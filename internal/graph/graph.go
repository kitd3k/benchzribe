package graph

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// Theme represents the graph theme
type Theme string

const (
	ThemeLight Theme = "light"
	ThemeDark  Theme = "dark"
)

// GenerateGraph creates a line chart from benchmark data
func GenerateGraph(data map[string][]float64) error {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Benchmark Results",
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show:    opts.Bool(true),
			Trigger: "axis",
		}),
		charts.WithLegendOpts(opts.Legend{
			Show: opts.Bool(true),
		}),
	)

	// Add X axis - use the first data series to determine length
	var xAxis []string
	var dataLen int
	for _, values := range data {
		dataLen = len(values)
		break
	}
	for i := 0; i < dataLen; i++ {
		xAxis = append(xAxis, fmt.Sprintf("Benchmark %d", i+1))
	}
	line.SetXAxis(xAxis)

	// Add series
	for name, values := range data {
		line.AddSeries(name, generateLineItems(values))
	}

	// Save the chart
	f, err := os.Create("benchmark-graph.html")
	if err != nil {
		return err
	}
	defer f.Close()

	return line.Render(f)
}

func generateLineItems(values []float64) []opts.LineData {
	items := make([]opts.LineData, 0, len(values))
	for _, v := range values {
		items = append(items, opts.LineData{Value: v})
	}
	return items
}

// GenerateMarkdownChart creates a simple text-based chart for embedding in README
func GenerateMarkdownChart(data map[string][]float64, benchmarkNames []string) string {
	var result strings.Builder
	
	// Create a simple bar chart using Unicode characters
	result.WriteString("```\n")
	result.WriteString("Performance Overview:\n")
	result.WriteString("====================\n\n")
	
	// Find max values for scaling
	maxNs := findMax(data["ns/op"])
	maxB := findMax(data["B/op"])
	maxAllocs := findMax(data["allocs/op"])
	
	for i, name := range benchmarkNames {
		if i >= len(data["ns/op"]) {
			continue
		}
		
		ns := data["ns/op"][i]
		b := data["B/op"][i]
		allocs := data["allocs/op"][i]
		
		result.WriteString(fmt.Sprintf("%-25s ", name))
		
		// Create simple bar visualization
		nsBar := createBar(ns, maxNs, 20)
		result.WriteString(fmt.Sprintf("ns/op: %s (%.0f)\n", nsBar, ns))
		
		if maxB > 0 {
			result.WriteString(fmt.Sprintf("%-25s ", ""))
			bBar := createBar(b, maxB, 20)
			result.WriteString(fmt.Sprintf("B/op:  %s (%.0f)\n", bBar, b))
		}
		
		if maxAllocs > 0 {
			result.WriteString(fmt.Sprintf("%-25s ", ""))
			allocsBar := createBar(allocs, maxAllocs, 20)
			result.WriteString(fmt.Sprintf("alloc: %s (%.0f)\n", allocsBar, allocs))
		}
		
		result.WriteString("\n")
	}
	result.WriteString("```\n")
	return result.String()
}

// GenerateMermaidChart creates a mermaid chart for GitHub display
func GenerateMermaidChart(data map[string][]float64, benchmarkNames []string) string {
	var result strings.Builder
	
	result.WriteString("```mermaid\n")
	result.WriteString("xychart-beta\n")
	result.WriteString("    title \"Benchmark Performance (ns/op)\"\n")
	result.WriteString("    x-axis [")
	
	// Add benchmark names (shortened)
	for i, name := range benchmarkNames {
		shortName := name
		if len(name) > 10 {
			shortName = name[:10]
		}
		if i > 0 {
			result.WriteString(", ")
		}
		result.WriteString(fmt.Sprintf("\"%s\"", shortName))
	}
	result.WriteString("]\n")
	
	result.WriteString("    y-axis \"Nanoseconds per Operation\"\n")
	result.WriteString("    line [")
	
	// Add ns/op values
	for i, val := range data["ns/op"] {
		if i > 0 {
			result.WriteString(", ")
		}
		result.WriteString(fmt.Sprintf("%.0f", val))
	}
	result.WriteString("]\n")
	result.WriteString("```\n")
	
	return result.String()
}

func findMax(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	max := values[0]
	for _, v := range values[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func createBar(value, max float64, width int) string {
	if max == 0 {
		return strings.Repeat("▁", width)
	}
	
	ratio := value / max
	filledWidth := int(ratio * float64(width))
	
	bar := strings.Repeat("█", filledWidth)
	bar += strings.Repeat("▁", width-filledWidth)
	
	return bar
}
