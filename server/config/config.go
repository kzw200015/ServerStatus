package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var config Config

type Config struct {
	Port  int          `yaml:"port"`
	Nodes []NodeConfig `yaml:"nodes"`
}

type NodeConfig struct {
	Id    string `yaml:"id"`
	Token string `yaml:"token"`
}

func Get() Config {
	return config
}

func GetAuthMap() map[string]string {
	authMap := make(map[string]string)
	for _, node := range config.Nodes {
		authMap[node.Id] = node.Token
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
