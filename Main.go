package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
	} `yaml:"database"`
}

const (
	defaultFilePath = "./config.yaml"
)

func main() {

	var filePath string
	if len(os.Args) <= 1 {
		filePath = defaultFilePath
	} else {
		filePath = os.Args[1]
	}

	config := parseConfig(filePath)

	log.Println("Server.Host:", config.Server.Host)
}

func parseConfig(filePath string) *Config {
	log.Println("Open", filePath)
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	log.Println("Parse", filePath)
	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Done", filePath)
	return &cfg
}
