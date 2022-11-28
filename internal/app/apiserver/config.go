package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/JadesHeart/http-rest-api/internal/app/store"
	"os"
)

// Config struct
type Config struct {
	BindAddr string `json:"bind_addr"`
	LogLevel string `json:"log_level"`
	Store    *store.Config
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

// NewConfig return new config
func NewConfig() *Config {
	configuration := loadConfiguration("C:/Users/лали/Desktop/проекты/http-rest-api/configs/apiserver.json")
	return &Config{
		BindAddr: configuration.BindAddr,
		LogLevel: configuration.LogLevel,
		Store:    store.NewConfig(),
	}
}
