package common_commands

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

func InitCommands() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.CompletionOptions.DisableNoDescFlag = true
	rootCmd.CompletionOptions.DisableDescriptions = true

	executeCommands()
}

func executeCommands() {
	rootCmd.Execute()
}
