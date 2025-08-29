package game

import (
	"fmt"
	"hangman/common"
	"slices"
	"strings"
	"unicode/utf8"
)

type Game struct {
	word       string
	resultGame string
	level      string
}

func NewGame() *Game {
	return new(Game)
}

func (g *Game) StartGameLoop() {
	playWordString := ""
	for i := 0; i < utf8.RuneCountInString(g.word); i++ {
		playWordString += string('-')
	}

	word := Split(g.word)
	wordFull := word
	playWord := Split(playWordString)
	wrongLetters := ""
	countOfAttempts := ReturnNumberOfAttempts(g.level)
	g.resultGame = ""

	fmt.Println()
	common.PrintYellowText("Игра началась ! Чтобы выйти в главное меню, введите EXIT")

	for {
		if countOfAttempts == 0 {
			g.resultGame = common.LOSE
			break
		} else if len(word) == 0 {
			g.resultGame = common.WIN
			break
		}

		fmt.Println()
		displayProgress(countOfAttempts, wrongLetters, playWord)

		inputLetter := g.ReadSingleChar()
		if g.resultGame == "EXIT" {
			break
		}

		if IsRussianLetter(inputLetter) && slices.Contains(word, string(inputLetter)) {
			fmt.Println("Такая буква есть !")
			for _, v := range FindAllIndexes(wordFull, string(inputLetter)) {
				playWord[v] = wordFull[v]
			}
			RemoveString(&word, string(inputLetter))
		} else {
			if !IsRussianLetter(inputLetter) {
				fmt.Println("Необходимо ввести русскую букву !")
				continue
			}

			if slices.Contains(playWord, string(inputLetter)) || strings.Contains(wrongLetters, string(inputLetter)) {
				fmt.Println("Буква уже была !")
				continue
			}

			if !slices.Contains(word, string(inputLetter)) {
				fmt.Println("Такой буквы нет !")
				countOfAttempts--
				wrongLetters += string(inputLetter)
				continue
			}
		}

	}
}

func (g *Game) DisplayResult() {
	switch g.resultGame {
	case "WIN":
		common.PrintYellowText("Слово: " + g.word + "\nВы выиграли ! Вы отгадали все буквы !\n")
	case "LOSE":
		common.PrintRedBackgroundText("Вы проиграли ! Загаданное слово: " + g.word + "\n")
	case "EXIT":
		common.PrintRedText("Вы вышли !")
	}
}

func (g *Game) ReadSingleChar() rune {
	var input string
	for {
		fmt.Scanln(&input)
		if strings.ToUpper(input) == "EXIT" {
			g.resultGame = strings.ToUpper(input)
			return 0
		} else if utf8.RuneCountInString(input) > 1 || input == "" {
			common.PrintRedText("Введите 1 букву !")
			continue
		} else {
			s := []rune(input)
			return rune(s[0])
		}
	}
}

func (g *Game) SetWord(word string) {
	g.word = word
}

func (g *Game) SetCountOfAttempts(level string) {
	g.level = level
}

func (g *Game) GetResultGame() string {
	return g.resultGame
}
