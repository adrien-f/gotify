package config

import (
	"bufio"
	"fmt"
	"github.com/howeyc/gopass"
	"log"
	"os"
	"strings"
)

func getUsername(configuration *Configuration) {
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

func getPassword(configuration *Configuration) {
	for {
		fmt.Print("Spotify Password (will not be stored): ")
		password := string(gopass.GetPasswd())
		if len(password) > 0 {
			configuration.Password = strings.TrimSpace(password)
			return
		} else {
			fmt.Println("Empty password, please try again")
		}
	}
}

func StartWizard(configuration *Configuration) *Configuration {
	getUsername(configuration)
	getPassword(configuration)
	return configuration
}
