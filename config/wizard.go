package config

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func GetUsername(configuration *Configuration) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Spotify Username: ")
		username, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if len(username)-1 > 0 {
			configuration.Username = strings.TrimSpace(username)
			return
		} else {
			fmt.Println("Empty username, please try again")
		}
	}
}

func GetPassword(configuration *Configuration) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Spotify Password (will not be stored): ")
		password, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if len(password)-1 > 0 {
			configuration.Password = strings.TrimSpace(password)
			return
		} else {
			fmt.Println("Empty password, please try again")
		}
	}
}

func StartWizard(configuration *Configuration) *Configuration {
	GetUsername(configuration)
	GetPassword(configuration)
	return configuration
}
