package main

import (
	"hangman/core"
	"hangman/service/command"
	"hangman/service/display"
	"hangman/service/game"
	"hangman/service/loader"
)

func CreateDI() *core.Core {
	loaderWord := loader.NewLoader("../resource/words.txt")

	gameLoop := game.NewGame()

	display := display.NewDisplay("../resource/help.txt", "Приветствую в игре Виселица !", "ГЛАВНОЕ МЕНЮ")

	command := command.NewCommand()

	core := core.NewCore(display, command, gameLoop, loaderWord)

	return core
}
