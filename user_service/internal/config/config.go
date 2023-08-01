package config

import (
	"gopkg.in/yaml.v3"

	"log"
	"os"
)

type (
	Config struct {
		DB   DB   `yaml:"db"`
		HTTP HTTP `yaml:"http"`
	}

	DB struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBname   string `yaml:"db-name"`
	}

	HTTP struct {
		Port int `yaml:"port"`
	}
)

func ParseConfig() (Config, error) {
	file, err := os.ReadFile("./cmd/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var cfg Config

	if err := yaml.Unmarshal(file, &cfg); err != nil {
		log.Fatal(err)
	}

	return cfg, nil
}
