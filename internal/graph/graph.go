package graph

import (
	"fmt"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/kitd3k/benchzribe/internal/parser"
)

// Theme represents the graph theme
type Theme string

const (
	ThemeLight Theme = "light"
	ThemeDark  Theme = "dark"
)

// GenerateGraph creates a benchmark visualization
func GenerateGraph(results []parser.Result, outputFile string, theme Theme) error {
	if len(results) == 0 {
		return fmt.Errorf("no results to graph")
	}

	// Create a new bar chart
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Benchmark Results",
			Left:  "center",
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: string(theme),
		}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
		charts.WithLegendOpts(opts.Legend{Show: true, Right: "10%"}),
	)

	// Prepare data
	names := make([]string, 0, len(results))
	nsOps := make([]opts.BarData, 0, len(results))
	bytesOps := make([]opts.BarData, 0, len(results))
	allocsOps := make([]opts.BarData, 0, len(results))

	for _, r := range results {
		names = append(names, r.Name)
		nsOps = append(nsOps, opts.BarData{Value: r.NsPerOp})
		bytesOps = append(bytesOps, opts.BarData{Value: r.BytesPerOp})
		allocsOps = append(allocsOps, opts.BarData{Value: r.AllocsPerOp})
	}

	// Add data to chart
	bar.SetXAxis(names).
		AddSeries("ns/op", nsOps).
		AddSeries("B/op", bytesOps).
		AddSeries("allocs/op", allocsOps)

	// Create output file
	f, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer f.Close()

	// Render the chart
	if err := bar.Render(f); err != nil {
		return fmt.Errorf("failed to render chart: %w", err)
	}

	return nil
}
