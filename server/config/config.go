package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var config Config

type Config struct {
	Port  int          `yaml:"port"`
	Token string       `yaml:"token"`
	Nodes []NodeConfig `yaml:"nodes"`
}

type NodeConfig struct {
	Id string `yaml:"id"`
}

func Get() Config {
	return config
}

func GetAuthMap() map[string]string {
	authMap := make(map[string]string)
	for _, node := range config.Nodes {
		authMap[node.Id] = config.Token
	}
	return authMap
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
