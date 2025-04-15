package main

import (
	"fmt"
	"os"
	"strings"
	"time"
  
	"github.com/kitd3k/benchzribe/internal/parser"
	"github.com/kitd3k/benchzribe/internal/readme"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: benchscribe <run|graph|readme>")
		return
	}
	cmd := os.Args[1]

	switch cmd {
	case "run":
		fmt.Println("🔍 Parsing benchmark results...")

		results, err := parser.Parse("bench.out")
		if err != nil {
			fmt.Println("❌ Failed to parse:", err)
			return
		}

		if len(results) == 0 {
			fmt.Println("⚠️ No benchmark results found.")
			return
		}

		var sb strings.Builder
		sb.WriteString("### 📊 Benchmark Results\n\n")
		sb.WriteString("| Benchmark | ns/op | B/op | allocs/op |\n")
		sb.WriteString("|-----------|-------|------|------------|\n")
		sb.WriteString("\n_Last updated: " + time.Now().Format(time.RFC1123) + "_\n")

		for _, r := range results {
			sb.WriteString(fmt.Sprintf("| %s | %.0f | %d | %d |\n", r.Name, r.NsPerOp, r.BytesPerOp, r.AllocsPerOp))
		}

		if err := readme.Update("README.md", sb.String()); err != nil {
			fmt.Println("❌ Failed to update README:", err)
			return
		}

		fmt.Println("✅ README updated with benchmark results!")

	case "graph":
		fmt.Println("📊 Graph support coming soon...")

	case "readme":
		fmt.Println("📝 Manual README update mode... (not used here)")

	default:
		fmt.Println("Unknown command:", cmd)
	}
}
