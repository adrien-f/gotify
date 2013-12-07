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

func getRemember(configuration *Configuration) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Do you want gotify to login you automatically ? [Y/n] ")
		remember, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if strings.TrimSpace(remember) == "" {
			configuration.RememberMe = true
			return
		} else if stringInSlice(strings.TrimSpace(remember), []string{"y", "Y", "yes", "Yes"}) {
			configuration.RememberMe = true
			return
		} else if stringInSlice(strings.TrimSpace(remember), []string{"n", "N", "no", "No"}) {
			configuration.RememberMe = false
			return
		} else {
			fmt.Println("I could not understand your answer, please write either yes or no")
		}
	}
}

// From http://stackoverflow.com/a/15323988
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func StartWizard(configuration *Configuration) *Configuration {
	getUsername(configuration)
	getPassword(configuration)
	getRemember(configuration)
	return configuration
}
