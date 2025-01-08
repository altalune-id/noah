package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port  int    `yaml:"port"`
		Mode  string `yaml:"mode"`
		Debug bool   `yaml:"debug"`
	} `yaml:"server"`

	Database struct {
		Driver string `yaml:"driver"`
		Uri    string `yaml:"uri"`
	} `yaml:"database"`
}

func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	config := &Config{}
	if err := decoder.Decode(config); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	return config, nil
}
