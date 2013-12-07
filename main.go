// This is gotify, a lightweight Spotify player written in Go
package main

import (
	"flag"
	"fmt"
	"github.com/adrien-f/gotify/config"
	sp "github.com/op/go-libspotify"
)

func Ascii() {
	fmt.Println("             _   _  __")
	fmt.Println("  __ _  ___ | |_(_)/ _|_   _")
	fmt.Println(" / _` |/ _ \\| __| | |_| | | |")
	fmt.Println("| (_| | (_) | |_| |  _| |_| |")
	fmt.Println(" \\__, |\\___/ \\__|_|_|  \\__, |")
	fmt.Println(" |___/                 |___/")
	fmt.Println("\ngotify version 1.2 - libspotify", sp.BuildId())
}

var (
	configuration *config.Configuration
	debug         = flag.Bool("debug", false, "debug output")
)

func main() {
	flag.Parse()
	Ascii()
	configuration = config.LoadConfig()
}
