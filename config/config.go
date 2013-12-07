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

// Starts the wizard config
func configWizard() *Configuration {
	configuration := new(Configuration)
	fmt.Println("Welcome to gotify !\nThis wizard will help you set up gotify, follow it carefully !")
	StartWizard(configuration)
	err := saveConfig(configuration)
	if err != nil {
		log.Fatal(err)
	}
	return configuration
}

// Save the configuration in the config.json file
func saveConfig(configuration *Configuration) error {
	config, err := json.Marshal(configuration)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("config.json", config, 0644)
	return err
}

// Load the configuration from config.json or launch the wizard if it does not exists
func LoadConfig() *Configuration {
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		configuration := configWizard()
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
