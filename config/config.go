package config

import (
	"os"
	"encoding/json"
	"log"
)

type Configuration struct {
	Firebase_url string
}

func FirebaseUrl() string{
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)

	if err != nil {
		log.Panic(err)
	}

	return configuration.Firebase_url
}




