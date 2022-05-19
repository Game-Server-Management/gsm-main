package commands

import (
	"gsm/config"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   filepath.Base(os.Args[0]),
	Short: config.Name,
}

func Init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.CompletionOptions.DisableNoDescFlag = true
	rootCmd.CompletionOptions.DisableDescriptions = true

	initConfigCommands()

	rootCmd.AddCommand(
		configCmd,
		installCmd,
		systemCmd,
		versionCmd,
	)

	executeCommands()
}

func executeCommands() {
	calculateVersion()

	rootCmd.SetVersionTemplate(display)
	rootCmd.Execute()
}
