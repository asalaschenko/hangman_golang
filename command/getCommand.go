package command

import (
	"fmt"
	"strconv"
)

const (
	MIN_LENGTH_WORD = 3
	MAX_LENGTH_WORD = 10
)

type Command struct {
	commandGame   string
	commandLevel  string
	commandLength int
}

func NewCommand() *Command {
	return new(Command)
}

func (c *Command) EnterLevelCommand() {
	var levelMap = map[string]string{
		"1": "EASY",
		"2": "MEDIUM",
		"3": "HARD",
	}

	fmt.Println("Введите 1, 2 или 3:")
	for {
		var text string
		fmt.Println("1-EASY\n2-MEDIUM\n3-HARD")
		fmt.Scanln(&text)

		if v, ok := levelMap[text]; ok {
			c.commandLevel = v
			break
		} else {
			fmt.Println("Введите 1, 2 или 3:")
		}
	}

	fmt.Println()
}

func (c *Command) EnterGameCommand() {
	var commandMap = map[string]string{
		"1": "START",
		"2": "HELP",
		"3": "EXIT",
	}

	fmt.Println("Введите 1, 2 или 3:")
	for {
		var text string
		fmt.Println("1-START\n2-HELP\n3-EXIT")
		fmt.Scanln(&text)
		if v, ok := commandMap[text]; ok {
			c.commandGame = v
			break
		} else {
			fmt.Println("Введите 1, 2 или 3:")
		}
	}

	fmt.Println()
}

func (c *Command) EnterLengthCommand() {
	fmt.Println("Введите длину слова от 3 до 8:")
	var value1 string
	for {
		fmt.Scanln(&value1)
		value2, err := strconv.Atoi(value1)
		if err == nil && value2 >= MIN_LENGTH_WORD && value2 <= MAX_LENGTH_WORD {
			c.commandLength = value2
			break
		} else {
			fmt.Println("Неверно ! Введите длину слова от 3 до 8:")
		}
	}
	fmt.Println()
}

func (c *Command) ReadCommandLength() int {
	return c.commandLength
}

func (c *Command) ReadCommandLevel() string {
	return c.commandLevel
}

func (c *Command) ReadCommandGame() string {
	return c.commandGame
}
