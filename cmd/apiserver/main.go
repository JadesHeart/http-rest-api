package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/JadesHeart/http-rest-api/internal/app/apiserver"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "C:/Users/лали/Desktop/проекты/jopa/http-rest-api/configs/apiserver.toml", "path to config file")
}
func main() {
	flag.Parse()

	config := apiserver.NewConfig() // Create new config
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	s := apiserver.New(config)        // Create new APIServer struct
	if err := s.Start(); err != nil { // Start API server
		log.Fatal(err)
	}
}
