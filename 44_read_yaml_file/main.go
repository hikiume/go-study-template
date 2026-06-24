package main

import (
	"fmt"
	"log"
	"os"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		User     string  `yaml:"user"`
		Password string  `yaml:"password"`
		Timeout  float64 `yaml:"timeout"`
	} `yaml:"database"`
	Tags []string `yaml:"tags"`
}

func main() {
	file, fileErr := os.ReadFile("example.yaml")
	if fileErr != nil {
		log.Fatalln("ファイルの読み込みに失敗")
	}

	var config Config

	err := yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalln("parseに失敗しました")
	}

	fmt.Printf("Server Port: %d\n", config.Server.Port)
	fmt.Printf("Server Host: %s\n", config.Server.Host)
	fmt.Printf("DB User: %s\n", config.Database.User)
	fmt.Printf("DB Timeout: %.1f\n", config.Database.Timeout)
	fmt.Printf("Tags: %v\n", config.Tags)
}
