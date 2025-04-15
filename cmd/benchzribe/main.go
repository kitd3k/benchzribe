package main

import (
	"fmt"
	"os"
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
		// Call parser.Parse("bench.out")
	case "graph":
		fmt.Println("📊 Generating graph...")
		// Call graph.Generate()
	case "readme":
		fmt.Println("📝 Updating README.md...")
		// Call readme.Update()
	default:
		fmt.Println("Unknown command:", cmd)
	}
}
