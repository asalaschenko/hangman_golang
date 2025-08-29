package game

import (
	"hangman/common"
	"strings"
)

func displayProgress(count int, wrongLetter string, playWord []string) {
	status := "Осталось жизней:"
	for range count {
		status += " |"
	}
	common.PrintRedText(status)

	status = "Неправильные буквы: "
	status += wrongLetter
	common.PrintRedText(status)

	common.PrintRedText(strings.Join(playWord, ""))
}
