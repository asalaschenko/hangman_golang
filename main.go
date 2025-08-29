package main

import (
	"hangman/core"
	"hangman/service/command"
	"hangman/service/display"
	"hangman/service/game"
	"hangman/service/loader"
)

func main() {
	game := сreateDI()
	game.StartGame()
}

func сreateDI() *core.Core {
	loaderWordObj := loader.NewLoader("./resource/words.txt")

	gameLoopObj := game.NewGame()

	startingPhrase := "Приветствую в игре Виселица !\n"
	textMenu := "ГЛАВНОЕ МЕНЮ"
	displayObj := display.NewDisplay("./resource/help.txt", startingPhrase, textMenu)

	commandObj := command.NewCommand()

	coreObj := core.NewCore(displayObj, commandObj, gameLoopObj, loaderWordObj)

	return coreObj
}
