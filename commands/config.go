package commands

import (
	"fmt"
	"gsm/commands/configuration"
	"gsm/config"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: fmt.Sprintf("For configuration of %s.", config.Name),
}

func initConfigCommands() {
	configCmd.CompletionOptions.DisableDefaultCmd = true
	configCmd.CompletionOptions.DisableNoDescFlag = true
	configCmd.CompletionOptions.DisableDescriptions = true

	configCmd.AddCommand(
		configuration.InitConfigCmd,
		configuration.UpdateConfigCmd,
	)

	configuration.InitConfigCmd.Flags().StringVar(&configuration.GamesPath, "games_path", "", "Path where games will be installed")
	configuration.UpdateConfigCmd.Flags().StringVar(&configuration.GamesPath, "games_path", "", "Path where games will be installed")
}
