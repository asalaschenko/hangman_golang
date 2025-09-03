package display

import "fmt"

const (
	RED            = "\u001B[31m"
	YELLOW         = "\u001B[33m"
	RED_BACKGROUND = "\u001B[41m"
	RESET          = "\u001B[0m"
)

func PrintYellowText(text string) {
	fmt.Println(YELLOW + text + RESET)
}

func PrintRedText(text string) {
	fmt.Println(RED + text + RESET)
}

func PrintRedBackgroundText(text string) {
	fmt.Println(RED_BACKGROUND + text + RESET)
}
