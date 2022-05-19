package main

import (
	"fmt"
	"gsm/config"
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
}
