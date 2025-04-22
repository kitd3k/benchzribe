# 🔧 Mock API

A command-line tool for parsing, visualizing, and managing Go benchmark results. Automatically updates your README with beautifully formatted benchmark results and provides visualization options.

## ✨ Features

- 📊 Parse Go benchmark results and generate formatted tables
- 📝 Automatically update README with benchmark results
- 📈 Generate visualizations of benchmark data (coming soon)
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

3. Generate performance graphs (coming soon):
```bash
benchzribe graph
```

## 📊 Benchmark Results

<!-- BENCHSCRIBE:START -->

| Benchmark | ns/op | B/op | allocs/op |
|-----------|-------|------|------------|

_Last updated: Tue, 22 Apr 2025 15:01:52 UTC_
| TestHandler-4 | 2081 | 6132 | 19 |

<!-- BENCHSCRIBE:END -->

## 🛠️ Commands

- `benchzribe run` - Parse benchmark results and update README
- `benchzribe graph` - Generate performance visualization (coming soon)
- `benchzribe readme` - Manually update README section

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.
