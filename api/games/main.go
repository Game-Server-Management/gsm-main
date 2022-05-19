package games

import (
	"fmt"
	"gsm/api"
	"gsm/utilities"
	"os"
	"path/filepath"
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

func InstallExecutable(game string) {
	executablePath := fmt.Sprintf("%s/%s", filepath.Dir(os.Args[0]), game)
	if utilities.FileExists(executablePath) {
		i := 1
		for utilities.FileExists(fmt.Sprintf("%s/%s-%d", filepath.Dir(os.Args[0]), game, i)) {
			i++
		}
		executablePath = fmt.Sprintf("%s-%d", executablePath, i)
	}

	err := utilities.DownloadFile(executablePath, fmt.Sprintf("%s/%s_%s_%s", api.ExecutablesURL, game, utilities.OS, utilities.Arch), 0755)
	if err != nil {
		panic(err)
	}

}
