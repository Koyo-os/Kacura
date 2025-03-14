package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct{
	Port uint `yaml:"manager_port"`
	Host string `yaml:"manager_port"`
	WorkerMaxCount int `yaml:"worker_max_count"`
	WorkerMaxProcent int `yaml:"worker_agent_procent_max_count"`
	WorkerPorts string `yaml:"worker_ports"`
}

func Init() (*Config, error) {
	file,err := os.Open("config.yaml")
	if err != nil{
		return nil, fmt.Errorf("cant load config: %v",err)
	}

	body,err := io.ReadAll(file)
	if err != nil{
		return nil, fmt.Errorf("cant read file: %v",err)
	}

	var cfg Config

	if err = yaml.Unmarshal(body, &cfg);err != nil{
		return nil, fmt.Errorf("cant unmarshal config: %v",err)
	}

	return &cfg, nil
}