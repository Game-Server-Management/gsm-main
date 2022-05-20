package main

import (
	"fmt"
	"gsm/config"
	common_commands "gsm/games/dev/common/commands"
	"gsm/utilities"
	"os"
	"path/filepath"

	"github.com/TwiN/go-color"
)

func main() {
	if !utilities.FileExists(config.ConfigPath) {
		fmt.Printf("%s(!)%s Configuration file doesn't exists. Use %s%s config init%s to create the configuration.\n", color.Red, color.Reset, color.Green, filepath.Base(os.Args[0]), color.Reset)
		return
	}

	config.Init()

	common_commands.InitCommands()
}
