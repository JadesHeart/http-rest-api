package store

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DatabaseURL string `json:"database_url"`
}

func loadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
func NewConfig() *Config {
	configuration := loadConfiguration("C:/Users/лали/Desktop/проекты/http-rest-api/configs/store.json")
	return &Config{
		DatabaseURL: configuration.DatabaseURL,
	}
}
