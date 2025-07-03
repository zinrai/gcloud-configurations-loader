package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Loads and parses the YAML configuration file
func LoadConfig(filepath string) (*ConfigFile, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var config ConfigFile
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
