package main

import (
	"github.com/JadesHeart/http-rest-api/internal/app/apiserver"
	"log"
)

func main() {
	config := apiserver.NewConfig()   // Create new config
	s := apiserver.New(config)        // Create new APIServer struct
	if err := s.Start(); err != nil { // Start API server
		log.Fatal(err)
	}

}
