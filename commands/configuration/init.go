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

var InitConfigCmd = &cobra.Command{
	Use:   "init",
	Short: fmt.Sprintf("To initiate the configuration of %s.", config.Name),
	Run:   initiateConfigExecuteCommand,
}

var GamesPath string

func initiateConfigExecuteCommand(cmd *cobra.Command, args []string) {
	if utilities.FileExists(config.ConfigPath) {
		fmt.Printf("%s(!)%s Configuration file already exists. Use %s%s config update%s to update the configuration.\n", color.Red, color.Reset, color.Green, filepath.Base(os.Args[0]), color.Reset)
		return
	}

	questions := make([]*survey.Question, 0)

	if GamesPath == "" {
		questions = append(questions, &survey.Question{
			Name: "games_path",
			Prompt: &survey.Input{
				Message: "Games Instalation Path:",
				Default: fmt.Sprintf("%s/games", filepath.Dir(os.Args[0])),
			},
			Validate: validateGamePath,
		})

		_ = survey.Ask(questions, &GamesPath)
	}

	config.Configuration.Games_path = GamesPath
	config.SaveConfig(config.ConfigPath)
	fmt.Printf("Configuration file has been succesfully saved. Use %s%s config update%s to update the configuration.\n", color.Green, filepath.Base(os.Args[0]), color.Reset)
}

func validateGamePath(val interface{}) error {
	path := val.(string)
	_, err := os.Stat(path)
	return err
}
