package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Configuration struct {
	Username   string `json:"username"`
	Password   string
	RememberMe bool `json:"remember_me"`
}

func ConfigWizard() *Configuration {
	configuration := new(Configuration)
	fmt.Println("Welcome to gotify !\nThis wizard will help you set up gotify, follow it carefully !")
	StartWizard(configuration)
	return configuration
}

func LoadConfig() *Configuration {
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		configuration := ConfigWizard()
		return configuration
	}
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	configuration := new(Configuration)
	err = json.Unmarshal(file, &configuration)
	if err != nil {
		log.Fatal(err)
	}
	return configuration
}
