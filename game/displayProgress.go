package game

import (
	"hangman/display"
	"strings"
)

func displayProgress(count int, wrongLetter string, playWord []string) {
	status := "Осталось жизней:"
	for range count {
		status += " |"
	}
	display.PrintRedText(status)

	status = "Неправильные буквы: "
	status += wrongLetter
	display.PrintRedText(status)

	display.PrintRedText(strings.Join(playWord, ""))
}
