package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var config *Config

type Config struct {
	Server struct {
		Port  int    `yaml:"port"`
		Env   string `yaml:"env"`
		Stage string `yaml:"stage"`
		Debug bool   `yaml:"debug"`
	} `yaml:"server"`

	WebSocket struct {
		Endpoint string `yaml:"endpoint"`
	} `yaml:"websocket"`

	AWS struct {
		Region string `yaml:"region"`
	} `yaml:"aws"`

	Database struct {
		Driver string `yaml:"driver"`
		Uri    string `yaml:"uri"`
	} `yaml:"database"`
}

func LoadConfig(filename string) (*Config, error) {
	if config != nil {
		return config, nil
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)

	config = &Config{}
	if err := decoder.Decode(config); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	return config, nil
}
