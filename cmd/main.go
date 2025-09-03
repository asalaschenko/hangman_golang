package main

import (
	"hangman/command"
	"hangman/display"
	"hangman/game"
	"hangman/loader"
)

func main() {

	display := display.NewDisplay("../resource/help.txt", "Приветствую в игре Виселица !", "ГЛАВНОЕ МЕНЮ")
	loader := loader.NewLoader("../resource/words.txt")
	command := command.NewCommand()

	game := game.NewGame(display, loader, command)

	game.Start()
}
