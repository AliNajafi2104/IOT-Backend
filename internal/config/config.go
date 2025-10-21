package config

import (
	"log"
	"os"

	yaml "gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	MQTT struct {
		Broker   string `yaml:"broker"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"mqtt"`
	DB struct {
		URI string `yaml:"uri"`
	} `yaml:"db"`
}

func NewConfig() (*Config, error) {
	f, err := os.Open("/home/ali/5-semester/internet-of-things/IOT-Backend/internal/config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("failed to close file %v", err)
		}
	}()

	cfg := &Config{}
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		log.Fatal(err)
	}
	return cfg, nil
}
