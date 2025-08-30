package core

import (
	"hangman/common"
)

type Core struct {
	DP  IDisplay
	GC  IGetCommand
	GL  IGameLoop
	LDR ILoader
}

func NewCore(DP IDisplay, GC IGetCommand, GL IGameLoop, LDR ILoader) *Core {
	value := &Core{
		DP:  DP,
		GC:  GC,
		GL:  GL,
		LDR: LDR}
	return value
}

func (c *Core) StartGame() {
	c.DP.DisplayStartPhrase()
	for {
		c.DP.DisplayMenu()
		c.GC.EnterGameCommand()
		switch c.GC.ReadCommandGame() {
		case "START":
			c.GC.EnterLevelCommand()
			valueLevel := c.GC.ReadCommandLevel()
			c.GL.SetCountOfAttempts(valueLevel)
			c.GC.EnterLengthCommand()
			valueLength := c.GC.ReadCommandLength()
			c.LDR.LoaderWord()
			word := c.LDR.SelectRandom(valueLength)
			c.GL.SetWord(word)
			c.GL.StartGameLoop()
			c.GL.DisplayResult()
		case "HELP":
			c.DP.DisplayHelp()
		case "EXIT":
			common.PrintRedText("GOOD BYE !")
			return
		}
	}
}
