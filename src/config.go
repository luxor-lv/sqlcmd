package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		Type             string `yaml:"type"`
		ConnectionString string `yaml:"url"`
	} `yaml:"database"`
}

func loadConfig() (*Config, error) {

	f, error := os.Open("./sqlc.yaml")
	if error != nil {
		homeDir, _ := os.UserHomeDir()
		f, error = os.Open(homeDir + "/sqlc.yaml")
		if error != nil {
			return nil, error
		}
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	error = decoder.Decode(&cfg)
	if error != nil {
		return nil, error
	}

	return &cfg, nil
}
