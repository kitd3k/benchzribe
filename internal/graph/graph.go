package graph

import (
	"fmt"
	"os"

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
