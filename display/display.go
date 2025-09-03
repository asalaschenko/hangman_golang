package display

import (
	"fmt"
	"log/slog"
	"os"
)

type Display struct {
	help           string
	StartingPhrase string
	TextMenu       string
}

func NewDisplay(helpInfoPath string, startintPhrase string, textMenu string) *Display {
	const op = "display.NewDisplay"
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError})).With(
		slog.String("op", op),
	)

	value := new(Display)
	data, err := os.ReadFile(helpInfoPath)
	if err != nil {
		logger.Error("error of reading file")
	}
	value.help = string(data)
	value.StartingPhrase = startintPhrase
	value.TextMenu = textMenu
	return value
}

func (d *Display) DisplayHelp() {
	PrintYellowText(d.help)
	fmt.Println()
}

func (d *Display) DisplayMenu() {
	PrintRedBackgroundText(d.TextMenu)
	fmt.Println()
}

func (d *Display) DisplayStartPhrase() {
	PrintYellowText(d.StartingPhrase)
	fmt.Println()
}
