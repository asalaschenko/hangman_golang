package main

import (
	"hangman/core"
	"hangman/service/command"
	"hangman/service/display"
	"hangman/service/game"
	"hangman/service/loader"
)

func CreateDI() *core.Core {
	loaderWordObj := loader.NewLoader("../resource/words.txt")

	gameLoopObj := game.NewGame()

	displayObj := display.NewDisplay("../resource/help.txt", "Приветствую в игре Виселица !", "ГЛАВНОЕ МЕНЮ")

	commandObj := command.NewCommand()

	coreObj := core.NewCore(displayObj, commandObj, gameLoopObj, loaderWordObj)

	return coreObj
}
