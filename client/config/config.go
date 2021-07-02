package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var config Config

type Config struct {
	ServerUrl string `yaml:"server_url"`
	Id        string `yaml:"id"`
	Token     string `yaml:"token"`
}

func Parse(path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Panicln(err)
	}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Panicln(err)
	}
}

func Get() Config {
	return config
}
