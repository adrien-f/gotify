package config

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func GetUsername(configuration *Configuration) {
	gotUsername := false
	for gotUsername == false {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Spotify Username: ")
		username, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if len(username)-1 > 0 {
			configuration.Username = strings.TrimSpace(username)
			gotUsername = true
		}
	}

}

func StartWizard(configuration *Configuration) *Configuration {
	GetUsername(configuration)
	return configuration
}
