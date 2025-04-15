#!/bin/bash

echo "🧪 Running mock API benchmark..."

cd mockapi || exit
go test ./internal/benchmarks -bench=. -benchmem > ../bench.out

cd ..
go run cmd/benchscribe/main.go run
go run cmd/benchscribe/main.go graph
go run cmd/benchscribe/main.go readme

echo "✅ Benchmark results parsed and README updated!"
