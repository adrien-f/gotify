// This is the playlists manager
package playlists

import (
	"fmt"
	"github.com/adrien-f/gotify/config"
	sp "github.com/op/go-libspotify"
	"strings"
)

var MenuCommands map[string]config.Command

func CreateCommands() {
	MenuCommands = map[string]config.Command{
		"help": cmdHelp,
		"ls":   listPlaylists,
	}
}

func cmdHelp(session *sp.Session, args []string, abort <-chan bool) error {
	for command := range MenuCommands {
		println(" - ", command)
	}
	return nil
}

func listPlaylists(session *sp.Session, args []string, abort <-chan bool) error {
	playlists, err := session.Playlists()
	if err != nil {
		return err
	}
	playlists.Wait()
	indent := 0
	for i := 0; i < playlists.Playlists() && i < 10; i++ {
		switch playlists.PlaylistType(i) {
		case sp.PlaylistTypeStartFolder:
			if folder, err := playlists.Folder(i); err == nil {
				fmt.Print(strings.Repeat(" ", indent))
				fmt.Println(folder.Name())
			}
			indent += 2
		case sp.PlaylistTypeEndFolder:
			indent -= 2
		case sp.PlaylistTypePlaylist:
			fmt.Print(strings.Repeat(" ", indent))
			fmt.Println("[", i, "]", playlistStr(playlists.Playlist(i)))
		}
	}
	return nil
}

func playlistStr(playlist *sp.Playlist) string {
	playlist.Wait()
	return fmt.Sprintf("♫ %s",
		playlist.Name(),
	)
}
