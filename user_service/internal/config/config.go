package config

import (
	"github.com/andReyM228/lib/database"
	"gopkg.in/yaml.v3"

	"log"
	"os"
)

type (
	Config struct {
		DB   database.DBConfig `yaml:"db"`
		HTTP HTTP              `yaml:"http"`
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
