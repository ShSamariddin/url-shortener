package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Env         string     `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string     `yaml:"storage_path" env-required:"true"`
	HTTPServer  HTTPServer `yaml:"http_server" env-required:"true"`
}

type HTTPServer struct {
	Address     string `yaml:"address" env-required:"true"`
	Timeout     string `yaml:"timeout" env-required:"true"`
	IdleTimeout string `yaml:"idle_timeout" env-required:"true"`
}

func MustLoad(fileName string) Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	var config Config
	err = yaml.Unmarshal(data, config)
	if err != nil {
		log.Fatal("Can't unmarshal config")
	}
	return config
}
