#!/bin/bash

echo "🔧 Testing CI Flow for BenchZribe"
echo "=================================="

echo "1. Building benchzribe..."
go build -v -o benchzribe ./cmd/benchzribe

echo "2. Running benchmarks..."
cd mockapi
go test ./internal/benchmarks -bench=. -benchmem > ../bench.out
cd ..

echo "3. Generated benchmark results:"
cat bench.out

echo "4. Updating README with benchzribe..."
./benchzribe run

echo "5. Generating graph..."
./benchzribe graph

echo "6. Checking git status..."
git status

echo "7. Files generated:"
ls -la *.html *.out 2>/dev/null || echo "No graph/bench files found"

echo ""
echo "✅ CI Flow test completed!"
echo "The README should now contain updated benchmark results."
echo "The benchmark-graph.html should contain an interactive chart."
