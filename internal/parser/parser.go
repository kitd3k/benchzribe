package parser

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

type BenchResult struct {
	Name        string
	NsPerOp     float64
	BytesPerOp  int
	AllocsPerOp int
}

func Parse(path string) ([]BenchResult, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var results []BenchResult
	re := regexp.MustCompile(`^Benchmark(\S+)\s+\d+\s+([\d.]+)\sns/op\s+(\d+)\sB/op\s+(\d+)\sallocs/op`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		m := re.FindStringSubmatch(line)
		if len(m) == 5 {
			ns, _ := strconv.ParseFloat(m[2], 64)
			bytes, _ := strconv.Atoi(m[3])
			allocs, _ := strconv.Atoi(m[4])
			results = append(results, BenchResult{
				Name:        m[1],
				NsPerOp:     ns,
				BytesPerOp:  bytes,
				AllocsPerOp: allocs,
			})
		}
	}
	return results, nil
}
