package display

import (
	"hangman/common"
	"os"
)

type Display struct {
	help           string
	StartingPhrase string
	TextMenu       string
}

func NewDisplay(helpInfoPath string, startintPhrase string, textMenu string) *Display {
	value := new(Display)
	data, err := os.ReadFile(helpInfoPath)
	common.CheckErr(err)
	value.help = string(data)
	value.StartingPhrase = startintPhrase
	value.TextMenu = textMenu
	return value
}

func (d *Display) DisplayHelp() {
	common.PrintYellowText(d.help)
}

func (d *Display) DisplayMenu() {
	common.PrintYellowText(d.StartingPhrase)
	common.PrintRedBackgroundText(d.TextMenu)
}
