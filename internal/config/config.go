package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Server  ServerConfig  `json:"server"`
	Logging LoggingConfig `json:"logging"`
}

type ServerConfig struct {
	Port int `json:"port"`
}

type LoggingConfig struct {
	Level    string `json:"level"`
	Filename string `json:"filename"`
}

func InitConfig() Config {
	configFile, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	var config Config
	decoder := json.NewDecoder(configFile)
	if err := decoder.Decode(&config); err != nil {
		panic(err)
	}

	return config
}
