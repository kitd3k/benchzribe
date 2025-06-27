package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/kitd3k/benchzribe/internal/config"
	"github.com/kitd3k/benchzribe/internal/graph"
	"github.com/kitd3k/benchzribe/internal/parser"
	"github.com/kitd3k/benchzribe/internal/readme"
)

const (
	defaultBenchFile = "bench.out"
	defaultReadme    = "README.md"
)

var (
	errInvalidCommand = errors.New("invalid command")
	errNoResults     = errors.New("no benchmark results found")
)

func init() {
	log.SetPrefix("benchzribe: ")
	log.SetFlags(log.Ldate | log.Ltime)
}

func validateInputFile(filename string) error {
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("benchmark file %q does not exist", filename)
		}
		return fmt.Errorf("error accessing benchmark file: %w", err)
	}
	return nil
}

func main() {
	// Parse command line flags
	runCmd := flag.NewFlagSet("run", flag.ExitOnError)
	graphCmd := flag.NewFlagSet("graph", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected 'run' or 'graph' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "run":
		runCmd.Parse(os.Args[2:])
		if err := run(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "graph":
		graphCmd.Parse(os.Args[2:])
		if err := generateGraph(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Printf("unknown subcommand: %s\n", os.Args[1])
		os.Exit(1)
	}
}

func run() error {
	cfg := config.DefaultConfig()
	return handleRun(cfg)
}

func generateGraph() error {
	cfg := config.DefaultConfig()
	return handleGraph(cfg)
}

func handleRun(cfg config.Config) error {
	log.Println("🔍 Parsing benchmark results...")

	if err := validateInputFile(cfg.BenchmarkFile); err != nil {
		return err
	}

	results, err := parser.Parse(cfg.BenchmarkFile)
	if err != nil {
		return fmt.Errorf("failed to parse benchmark results: %w", err)
	}

	if len(results) == 0 {
		return errNoResults
	}

	var sb strings.Builder
	sb.WriteString("### 📊 Benchmark Results\n\n")
	sb.WriteString("| Benchmark | ns/op | B/op | allocs/op |\n")
	sb.WriteString("|-----------|-------|------|------------|\n")

	for _, r := range results {
		sb.WriteString(fmt.Sprintf("| %s | %.0f | %d | %d |\n", r.Name, r.NsPerOp, r.BytesPerOp, r.AllocsPerOp))
	}

	sb.WriteString("\n_Last updated: " + time.Now().Format(time.RFC1123) + "_\n")

	if err := readme.Update(cfg.ReadmeFile, sb.String()); err != nil {
		return fmt.Errorf("failed to update README: %w", err)
	}

	log.Println("✅ README updated with benchmark results!")

	// Generate graph if enabled
	if err := handleGraph(cfg); err != nil {
		log.Printf("⚠️ Warning: failed to generate graph: %v\n", err)
	}

	return nil
}

func handleGraph(cfg config.Config) error {
	log.Println("📊 Generating benchmark visualization...")

	results, err := parser.Parse(cfg.BenchmarkFile)
	if err != nil {
		return fmt.Errorf("failed to parse benchmark results: %w", err)
	}

	// Convert results to graph data format
	data := make(map[string][]float64)
	for _, r := range results {
		data["ns/op"] = append(data["ns/op"], float64(r.NsPerOp))
		data["B/op"] = append(data["B/op"], float64(r.BytesPerOp))
		data["allocs/op"] = append(data["allocs/op"], float64(r.AllocsPerOp))
	}

	// Generate graph
	if err := graph.GenerateGraph(data); err != nil {
		return fmt.Errorf("failed to generate graph: %w", err)
	}

	log.Printf("✅ Graph generated at %s\n", cfg.GraphOutput)
	return nil
}

func handleReadme(cfg config.Config) error {
	log.Println("📝 Manual README update mode...")
	return handleRun(cfg)
}

func handleConfig(cfg config.Config) error {
	log.Printf("Current configuration:\n")
	log.Printf("  Benchmark file: %s\n", cfg.BenchmarkFile)
	log.Printf("  README file: %s\n", cfg.ReadmeFile)
	log.Printf("  Graph output: %s\n", cfg.GraphOutput)
	log.Printf("  Theme: %s\n", cfg.Theme)
	return nil
}

func printUsage() {
	fmt.Printf(`Usage: %s <command>

Commands:
  run     Parse benchmark results and update README
  graph   Generate performance visualization
  readme  Manual README update
  config  Show current configuration

Example:
  %s run    # Parse bench.out and update README.md
`, filepath.Base(os.Args[0]), filepath.Base(os.Args[0]))
}
