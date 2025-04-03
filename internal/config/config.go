package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port   uint         `yaml:"manager_port"`
	Host   string       `yaml:"manager_port"`
	Worker WorkerConfig `yaml:"worker"`
	Smpt   SmptConfig   `yaml:"smpt"`
}

type SmptConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type WorkerConfig struct {
	MaxCount          int    `yaml:"max"`
	MaxProcentFromReq int    `yaml:"agent_procent"`
	Ports             string `yaml:"ports"`
}

func Init() (*Config, error) {
	file, err := os.Open("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("cant load config: %v", err)
	}

	var cfg Config

	if err = yaml.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshal config: %v", err)
	}

	return &cfg, nil
}
