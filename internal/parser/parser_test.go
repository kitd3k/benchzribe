package parser

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantErr  bool
		wantLen  int
		wantName string
		wantNs   float64
		wantB    int64
		wantA    int64
	}{
		{
			name: "valid benchmark result",
			input: `goos: linux
goarch: amd64
pkg: github.com/kitd3k/benchzribe
BenchmarkTest-8   	     100	  10000000 ns/op	   12345 B/op	      42 allocs/op
`,
			wantErr:  false,
			wantLen:  1,
			wantName: "Test-8",
			wantNs:   10000000,
			wantB:    12345,
			wantA:    42,
		},
		{
			name: "multiple benchmark results",
			input: `goos: linux
goarch: amd64
pkg: github.com/kitd3k/benchzribe
BenchmarkTest1-8   	     100	  10000000 ns/op	   12345 B/op	      42 allocs/op
BenchmarkTest2-8   	     200	   5000000 ns/op	    6789 B/op	      21 allocs/op
`,
			wantErr: false,
			wantLen: 2,
		},
		{
			name:    "empty file",
			input:   "",
			wantErr: false,
			wantLen: 0,
		},
		{
			name: "invalid format",
			input: `goos: linux
goarch: amd64
pkg: github.com/kitd3k/benchzribe
BenchmarkInvalid   	invalid format
`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary file
			tmpFile := filepath.Join(t.TempDir(), "bench.out")
			if err := os.WriteFile(tmpFile, []byte(tt.input), 0644); err != nil {
				t.Fatalf("failed to create test file: %v", err)
			}

			// Parse the file
			results, err := Parse(tmpFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if got := len(results); got != tt.wantLen {
				t.Errorf("Parse() got %d results, want %d", got, tt.wantLen)
				return
			}

			if tt.wantLen > 0 && tt.wantName != "" {
				if got := results[0].Name; got != tt.wantName {
					t.Errorf("Parse() first result name = %q, want %q", got, tt.wantName)
				}
				if got := results[0].NsPerOp; got != tt.wantNs {
					t.Errorf("Parse() first result ns/op = %f, want %f", got, tt.wantNs)
				}
				if got := results[0].BytesPerOp; got != tt.wantB {
					t.Errorf("Parse() first result B/op = %d, want %d", got, tt.wantB)
				}
				if got := results[0].AllocsPerOp; got != tt.wantA {
					t.Errorf("Parse() first result allocs/op = %d, want %d", got, tt.wantA)
				}
			}
		})
	}
} 