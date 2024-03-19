package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug" env-required:"true"`
	Listen  struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIP string `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"listen"`
	Storage StorageConfig `yaml:"storage"`
}

type StorageConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetConfig(path string) *Config {
	c := &Config{}

	yamlFile, err := os.ReadFile(path) // to rework
	if err != nil {
		log.Fatalf("config.GetConfig err #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal yaml config: %v", err)
	}

	log.Printf("Loaded config from %s", path)

	return c
}
