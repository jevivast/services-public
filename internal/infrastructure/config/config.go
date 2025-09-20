package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config holds all configuration for the application
type Config struct {
	Server struct {
		Port  string `yaml:"port"`
		Env   string `yaml:"env"`
		Limiter struct {
			Enabled bool `yaml:"enabled"`
			RPS     int  `yaml:"rps"`
			Burst   int  `yaml:"burst"`
		} `yaml:"limiter"`
	} `yaml:"server"`
}

// LoadConfig reads configuration from file or environment variables
func LoadConfig(path string) (*Config, error) {
	config := &Config{}

	// Read config file
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	// Init new YAML decoder
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	return config, nil
}

// GetConfigPath gets the path to the config file
func GetConfigPath(configPath string) (string, error) {
	if configPath == "" {
		configPath = os.Getenv("CONFIG_PATH")
		if configPath == "" {
			getwd, err := os.Getwd()
			if err != nil {
				return "", fmt.Errorf("error getting current working directory: %w", err)
			}
			configPath = filepath.Join(getwd, "configs", "config.yaml")
		}
	}

	return configPath, nil
}
