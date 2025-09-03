package game

import (
	"fmt"
	"hangman/display"
	"regexp"
	"slices"
	"strings"
	"unicode/utf8"
)

type GameLoop struct {
	word       string
	resultGame string
	level      string
}

var NumberOfAttemptsMap = map[string]int{
	"HARD":   5,
	"MEDIUM": 7,
	"EASY":   10,
}

// func NewGameLoop() *GameLoop {
// 	return new(GameLoop)
// }

func (g *GameLoop) StartGameLoop() {
	playWordString := ""
	for i := 0; i < utf8.RuneCountInString(g.word); i++ {
		playWordString += string('-')
	}

	word := Split(g.word)
	wordFull := word
	playWord := Split(playWordString)
	wrongLetters := ""
	countOfAttempts := NumberOfAttemptsMap[g.level]
	g.resultGame = ""

	display.PrintYellowText("Игра началась ! Чтобы выйти в главное меню, введите EXIT")
	fmt.Println()

	for {
		if countOfAttempts == 0 {
			g.resultGame = "LOSE"
			break
		} else if len(word) == 0 {
			g.resultGame = "WIN"
			break
		}

		displayProgress(countOfAttempts, wrongLetters, playWord)

		inputLetter := g.ReadSingleChar()
		if g.resultGame == "EXIT" {
			break
		}

		if IsRussianLetter(inputLetter) && slices.Contains(word, string(inputLetter)) {
			fmt.Println("Такая буква есть !")
			fmt.Println()
			for _, v := range FindAllIndexes(wordFull, string(inputLetter)) {
				playWord[v] = wordFull[v]
			}
			RemoveString(&word, string(inputLetter))
		} else {
			if !IsRussianLetter(inputLetter) {
				fmt.Println("Необходимо ввести русскую букву !")
				fmt.Println()
				continue
			}

			if slices.Contains(playWord, string(inputLetter)) || strings.Contains(wrongLetters, string(inputLetter)) {
				fmt.Println("Буква уже была !")
				fmt.Println()
				continue
			}

			if !slices.Contains(word, string(inputLetter)) {
				fmt.Println("Такой буквы нет !")
				fmt.Println()
				countOfAttempts--
				wrongLetters += string(inputLetter)
				continue
			}
		}
	}
}

func (g *GameLoop) DisplayResult() {
	switch g.resultGame {
	case "WIN":
		fmt.Println()
		display.PrintYellowText("Слово: " + strings.ToUpper(g.word) + "\nВы выиграли ! Вы отгадали все буквы !\n")
	case "LOSE":
		fmt.Println()
		display.PrintRedBackgroundText("Вы проиграли ! Загаданное слово: " + g.word)
		fmt.Println()
	case "EXIT":
		fmt.Println()
		display.PrintRedText("Вы вышли !\n")
	}
}

func (g *GameLoop) ReadSingleChar() rune {
	var input string
	for {
		fmt.Scanln(&input)
		if strings.ToUpper(input) == "EXIT" {
			g.resultGame = strings.ToUpper(input)
			return 0
		} else if utf8.RuneCountInString(input) > 1 || input == "" {
			display.PrintRedText("Введите 1 букву !")
			continue
		} else {
			s := []rune(input)
			return rune(s[0])
		}
	}
}

func IsRussianLetter(char rune) bool {
	// Регулярное выражение для русских букв (включая ё)
	russianPattern := regexp.MustCompile(`^[а-яёА-ЯЁ]$`)
	return russianPattern.MatchString(string(char))
}

func IsEnglishLetter(char rune) bool {
	// Регулярное выражение для английских букв
	englishPattern := regexp.MustCompile(`^[a-zA-Z]$`)
	return englishPattern.MatchString(string(char))
}

func Split(str string) []string {
	array := strings.Split(str, "")
	return array
}

func FindAllIndexes(slice []string, target string) []int {
	var indexes []int

	for i, value := range slice {
		if value == target {
			indexes = append(indexes, i)
		}
	}

	return indexes
}

func RemoveString(slice *[]string, target string) {
	var result []string

	for _, value := range *slice {
		if value != target {
			result = append(result, value)
		}
	}

	*slice = result
}
