// This is gotify, a lightweight Spotify player written in Go
package main

import (
	"flag"
	"fmt"
	"github.com/adrien-f/gotify/config"
	"github.com/adrien-f/gotify/playlists"
	"github.com/gobs/readline"
	sp "github.com/op/go-libspotify"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"time"
)

var (
	Configuration *config.Configuration
	debug         = flag.Bool("debug", false, "debug output")
	Session       *sp.Session
	reCommand     = regexp.MustCompile(`\s+`)
)

var menuCommands map[string]config.Command
var menuMatches []string

func Ascii() {
	fmt.Println("             _   _  __")
	fmt.Println("  __ _  ___ | |_(_)/ _|_   _")
	fmt.Println(" / _` |/ _ \\| __| | |_| | | |")
	fmt.Println("| (_| | (_) | |_| |  _| |_| |")
	fmt.Println(" \\__, |\\___/ \\__|_|_|  \\__, |")
	fmt.Println(" |___/                 |___/")
	fmt.Println("\ngotify version 1.2 - libspotify", sp.BuildId())
}

func createCommands() {
	menuCommands = map[string]config.Command{
		"playlists": cmdPlaylists,
	}
	menuMatches = make([]string, 0, len(menuCommands))
}

func AttemptedCompletion(text string, start, end int) []string {
	if start == 0 { // this is the command to match
		return readline.CompletionMatches(text, completionEntry)
	} else {
		return nil
	}
}

func completionEntry(prefix string, index int) string {
	if index == 0 {
		menuMatches = menuMatches[:0]
		for command := range menuCommands {
			if strings.HasPrefix(command, prefix) {
				menuMatches = append(menuMatches, command)
			}
		}
	}
	if index < len(menuMatches) {
		return menuMatches[index]
	} else {
		return ""
	}
}

func main() {
	flag.Parse()
	createCommands()
	Ascii()
	Configuration = config.LoadConfig()
	session, err := sp.NewSession(&sp.Config{
		ApplicationKey:   config.Spotify_key(),
		ApplicationName:  "Gotify",
		CacheLocation:    "tmp",
		SettingsLocation: "tmp",
	})
	if err != nil {
		log.Fatal(err)
	}
	_ = session

	exit := make(chan bool)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)
	go func() {
		for _ = range signals {
			select {
			case exit <- true:
			default:
			}
		}
	}()
	go func() {
		exitAttempts := 0
		running := true
		for running {
			if *debug {
				println("waiting for connection state change", session.ConnectionState())
			}

			select {
			case message := <-session.LogMessages():
				if *debug {
					println("!! log message", message.String())
				}
			case <-session.ConnectionStateUpdates():
				if *debug {
					println("!! connstate", session.ConnectionState())
				}
			case err := <-session.LoginUpdates():
				if *debug {
					fmt.Println(session.CurrentUser())
					println("!! login updated", err)
				}
			case <-session.LogoutUpdates():
				if *debug {
					println("!! logout updated")
				}
				running = false
			case _ = <-session.CredentialsBlobUpdates():
				if *debug {
					println("!! blob updated")
				}
			case <-exit:
				if *debug {
					println("!! exiting")
				}
				if exitAttempts >= 3 {
					os.Exit(42)
				}
				exitAttempts++
				session.Logout()
			case <-time.After(5 * time.Second):
				if *debug {
					println("state change timeout")
				}
			}
		}

		session.Close()
		os.Exit(32)
	}()

	if len(Configuration.Password) > 0 {
		credentials := sp.Credentials{
			Username: Configuration.Username,
			Password: Configuration.Password,
		}
		if err = session.Login(credentials, Configuration.RememberMe); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := session.Relogin(); err != nil {
			log.Fatal(err)
		}
	}
	mainLoop(session, "main", menuCommands)
}

func mainLoop(session *sp.Session, menu string, commands map[string]config.Command) {
	prompt := "gotify:[" + menu + "] > "
	readline.SetCompletionEntryFunction(completionEntry)
L:
	for {
		result := readline.ReadLine(&prompt)
		if result == nil {
			break L
		}
		line := *result

		switch line {
		case "":
			continue
		case "exit", "quit":
			break L
		case "help":
			cmdHelp(commands)
			continue
		default:
			args := reCommand.Split(line, -1)
			cmd := commands[args[0]]
			if cmd == nil && args[0] != "help" {
				fmt.Println("Unknown command:", args[0])
				cmdHelp(commands)
				continue
			}
			if err := cmd(session, args, nil); err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}

func cmdHelp(commands map[string]config.Command) error {
	for command := range commands {
		println(" - ", command)
	}
	return nil
}

func cmdPlaylists(session *sp.Session, args []string, abort <-chan bool) error {
	playlists.CreateCommands()
	mainLoop(session, "playlists", playlists.MenuCommands)
	return nil
}
