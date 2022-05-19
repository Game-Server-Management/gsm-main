package main

import (
	"gsm/api/games"
	"gsm/commands"
	"gsm/config"
)

func main() {
	config.Init()
	games.LoadGames()
	commands.Init()
}
