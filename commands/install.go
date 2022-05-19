package commands

import (
	"fmt"
	"gsm/api/games"
	"gsm/config"
	"gsm/utilities"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/TwiN/go-color"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install [game]",
	Short: fmt.Sprintf("To install a game using %s.", config.Name),
	Args:  cobra.MinimumNArgs(0),
	Run:   installExecuteCommand,
}

func installExecuteCommand(cmd *cobra.Command, args []string) {
	if !utilities.FileExists(config.ConfigPath) {
		fmt.Printf("%s(!)%s Configuration file doesn't exists. Use %s%s config init%s to create the configuration.\n", color.Red, color.Reset, color.Green, filepath.Base(os.Args[0]), color.Reset)
		return
	}

	game := ""
	if len(args) == 0 {
		_ = survey.AskOne(&survey.Select{
			Message: fmt.Sprintf("Please select a game to install using %s", config.Name),
			Options: games.DisplayGames,
		}, &game)

		splitGameName := strings.Split(game, "-")
		if len(splitGameName) < 2 {
			return
		}
		game = strings.Trim(splitGameName[1], " ")
	} else {
		game = args[0]
	}

	if !games.GameExists(game) {
		fmt.Printf("%s(!)%s The game doesn't exists. Please select one from %s%s install%s.\n", color.Red, color.Reset, color.Green, filepath.Base(os.Args[0]), color.Reset)
		return
	}

	games.InstallGame(game)
}
