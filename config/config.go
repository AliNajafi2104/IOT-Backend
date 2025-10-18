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
}

func LoadConfig() (*Config, error) {
	f, err := os.Open("/home/ali/5-semester/internet-of-things/IOT-Backend/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	cfg := &Config{}
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		log.Fatal(err)
	}
	return cfg, nil
}
