package game

import "hangman/display"

type Game struct {
	display    Display
	loader     Loader
	getCommand GetCommand
	gameLoop   GameLoop
}

func NewGame(display Display, loader Loader, getCommand GetCommand) *Game {
	return &Game{
		display:    display,
		loader:     loader,
		getCommand: getCommand,
	}
}

func (g *Game) Start() {
	g.display.DisplayStartPhrase()
	for {
		g.display.DisplayMenu()
		g.getCommand.EnterGameCommand()
		switch g.getCommand.ReadCommandGame() {
		case "START":
			g.getCommand.EnterLevelCommand()
			g.gameLoop.level = g.getCommand.ReadCommandLevel()
			g.getCommand.EnterLengthCommand()
			g.loader.LoadWord()
			g.gameLoop.word = g.loader.SelectRandom(g.getCommand.ReadCommandLength())
			g.gameLoop.StartGameLoop()
			g.gameLoop.DisplayResult()
		case "HELP":
			g.display.DisplayHelp()
		case "EXIT":
			display.PrintRedText("GOOD BYE !")
			return
		}
	}
}
