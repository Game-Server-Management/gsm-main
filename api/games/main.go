package games

import (
	"fmt"
	"gsm/api"
	"gsm/utilities"
	"strings"

	"github.com/TwiN/go-color"
)

var executableNames = make(map[string]bool, 0)
var gameNames = make(map[string]string, 0)
var DisplayGames = make([]string, 0)

func GameExists(executable string) bool {
	return executableNames[executable]
}

func LoadGames() {
	api.LoadAPI()
	games := api.Games[utilities.Arch]
	for game := range games {
		gameExecutable := strings.ToLower(strings.Trim(strings.Split(games[game], "-")[1], " "))
		executableNames[gameExecutable] = true
		gameNames[gameExecutable] = strings.Trim(strings.Split(games[game], "-")[0], " ")
	}

	for game := range executableNames {
		DisplayGames = append(DisplayGames, fmt.Sprintf("%s%s%s - %s", color.Green, gameNames[game], color.Reset, game))
	}
}

func InstallGame(game string) {

}
