package commands

import (
	"crypto/sha256"
	"fmt"
	"gsm/config"
	"io"
	"os"

	"github.com/TwiN/go-color"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: fmt.Sprintf("Shows up the version of %s.", config.Name),
	Run:   versionExecuteCommand,
}

var version = "1.0.0"
var hash = "none"
var display = "none"

func versionExecuteCommand(cmd *cobra.Command, args []string) {
	fmt.Printf("The version of %s is %s.\n", config.Name, display)
}

func calculateVersion() {
	f, err := os.Open(os.Args[0])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		panic(err)
	}

	hash = fmt.Sprintf("%x", h.Sum(nil))
	display = fmt.Sprintf("%sv%s%s (%ssha256:%s%s)", color.Green, version, color.Reset, color.Green, hash, color.Reset)
}
