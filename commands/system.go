package commands

import (
	"fmt"
	"gsm/utilities"
	"strings"

	"github.com/TwiN/go-color"
	"github.com/spf13/cobra"
)

var systemCmd = &cobra.Command{
	Use:   "system",
	Short: "Detects what OS are you running and what architecture.",
	Run:   systemExecuteCommand,
}

func systemExecuteCommand(cmd *cobra.Command, args []string) {
	fmt.Printf("The system is running %s%s %s%s.\n", color.Green, strings.Title(utilities.OS), strings.ToUpper(utilities.Arch), color.Reset)
}
