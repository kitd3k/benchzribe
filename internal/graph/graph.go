package graph

import (
	"os"

	"github.com/guptarohit/asciigraph"
)

func Generate(results []float64, path string) error {
	graph := asciigraph.Plot(results, asciigraph.Caption("Latency (ns/op)"))
	return os.WriteFile(path, []byte(graph), 0644)
}
