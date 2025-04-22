package parser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Result represents a single benchmark result
type Result struct {
	Name        string
	NsPerOp     float64
	BytesPerOp  int64
	AllocsPerOp int64
}

// Parse reads a Go benchmark output file and returns the parsed results
func Parse(path string) ([]Result, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var results []Result
	var hasBenchmark bool
	re := regexp.MustCompile(`^Benchmark([^\s]+)\s+\d+\s+([\d.]+)\s+ns/op\s+(\d+)\s+B/op\s+(\d+)\s+allocs/op`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Benchmark") {
			hasBenchmark = true
			m := re.FindStringSubmatch(line)
			if len(m) != 5 {
				return nil, fmt.Errorf("invalid benchmark format: %s", line)
			}

			ns, err := strconv.ParseFloat(m[2], 64)
			if err != nil {
				return nil, fmt.Errorf("invalid ns/op value: %s", m[2])
			}

			bytes, err := strconv.ParseInt(m[3], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid B/op value: %s", m[3])
			}

			allocs, err := strconv.ParseInt(m[4], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid allocs/op value: %s", m[4])
			}

			results = append(results, Result{
				Name:        m[1],
				NsPerOp:     ns,
				BytesPerOp:  bytes,
				AllocsPerOp: allocs,
			})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	// If we found a benchmark line but couldn't parse it, return error
	if hasBenchmark && len(results) == 0 {
		return nil, fmt.Errorf("no valid benchmark results found")
	}

	return results, nil
}
