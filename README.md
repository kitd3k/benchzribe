# 🔧 Mock API

A command-line tool for parsing, visualizing, and managing Go benchmark results. Automatically updates your README with beautifully formatted benchmark results and provides visualization options.

## ✨ Features

- 📊 Parse Go benchmark results and generate formatted tables
- 📝 Automatically update README with benchmark results
- 📈 Generate interactive visualizations of benchmark data
- ⏱️ Track performance changes over time
- 🔄 Easy integration with CI/CD pipelines

## 📦 Installation

```bash
go install github.com/kitd3k/benchzribe@latest
```

## 🚀 Usage

1. Run your Go benchmarks and save the output:
```bash
go test -bench . -benchmem ./... > bench.out
```

2. Update README with benchmark results:
```bash
benchzribe run
```

3. Generate interactive performance graphs:
```bash
benchzribe graph
```

## 📊 Benchmark Results

<!-- BENCHSCRIBE:START -->
### 📊 Benchmark Results

| Benchmark | ns/op | B/op | allocs/op |
|-----------|-------|------|------------|
| SimpleOperation-4 | 628 | 0 | 0 |
| StringConcatenation-4 | 5944 | 21080 | 99 |
| SliceOperations-4 | 734 | 0 | 0 |
| TestHandler-4 | 1958 | 5747 | 18 |

📈 **[View Interactive Graph](benchmark-graph.html)**

_Last updated: Fri, 27 Jun 2025 01:11:24 UTC_

<!-- BENCHSCRIBE:END -->

## 🛠️ Commands

- `benchzribe run` - Parse benchmark results and update README
- `benchzribe graph` - Generate interactive performance visualization
- `benchzribe readme` - Manually update README section

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.
