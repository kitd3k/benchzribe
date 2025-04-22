package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Config represents the application configuration
type Config struct {
	BenchmarkFile string `json:"benchmark_file"`
	ReadmeFile    string `json:"readme_file"`
	GraphOutput   string `json:"graph_output"`
	Theme         string `json:"theme"`
}

// DefaultConfig returns the default configuration
func DefaultConfig() Config {
	return Config{
		BenchmarkFile: "bench.out",
		ReadmeFile:    "README.md",
		GraphOutput:   "benchmark-graph.png",
		Theme:         "light",
	}
}

// LoadConfig loads the configuration from the specified file
func LoadConfig(filename string) (Config, error) {
	config := DefaultConfig()

	if filename == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return config, err
		}
		filename = filepath.Join(home, ".config", "benchzribe", "config.json")
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return config, nil
		}
		return config, err
	}

	if err := json.Unmarshal(data, &config); err != nil {
		return config, err
	}

	return config, nil
}

// SaveConfig saves the configuration to the specified file
func SaveConfig(config Config, filename string) error {
	if filename == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		configDir := filepath.Join(home, ".config", "benchzribe")
		if err := os.MkdirAll(configDir, 0755); err != nil {
			return err
		}
		filename = filepath.Join(configDir, "config.json")
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
} 