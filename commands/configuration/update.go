package configuration

import (
	"fmt"
	"gsm/config"
	"gsm/utilities"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/TwiN/go-color"
	"github.com/spf13/cobra"
)

var UpdateConfigCmd = &cobra.Command{
	Use:   "update",
	Short: fmt.Sprintf("To update the configuration of %s.", config.Name),
	Run:   updateConfigExecuteCommand,
}

func updateConfigExecuteCommand(cmd *cobra.Command, args []string) {
	if !utilities.FileExists(config.ConfigPath) {
		fmt.Printf("%s(!)%s Configuration file doesn't exists. Use %s%s config init%s to create the configuration.\n", color.Red, color.Reset, color.Green, filepath.Base(os.Args[0]), color.Reset)
		return
	}
	if GamesPath == "" {
		GamesPath = config.Configuration.Games_path

		questions := make([]*survey.Question, 0)
		questions = append(questions, &survey.Question{
			Name: "games_path",
			Prompt: &survey.Input{
				Message: "Games Instalation Path:",
				Default: GamesPath,
			},
			Validate: validateGamePath,
		})

		_ = survey.Ask(questions, &GamesPath)
	}

	config.Configuration.Games_path = GamesPath
	config.SaveConfig(config.ConfigPath)
	fmt.Print("Configuration file has been succesfully saved.\n")
}
