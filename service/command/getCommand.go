package command

import (
	"fmt"
	"strconv"
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
	fmt.Println("Введите 1, 2 или 3")
	for {
		flag := true
		var text string
		fmt.Println("1-EASY\n2-MEDIUM\n3-HARD")
		fmt.Scanln(&text)
		switch text {
		case "1":
			c.commandLevel = "EASY"
		case "2":
			c.commandLevel = "MEDIUM"
		case "3":
			c.commandLevel = "HARD"
		default:
			flag = false
			fmt.Println("Введите 1, 2 или 3")
		}
		if flag {
			break
		}
	}
}

func (c *Command) EnterGameCommand() {
	fmt.Println("Введите 1, 2 или 3")
	for {
		flag := true
		var text string
		fmt.Println("1-START\n2-HELP\n3-EXIT")
		fmt.Scanln(&text)
		switch text {
		case "1":
			c.commandLevel = "START"
		case "2":
			c.commandLevel = "HELP"
		case "3":
			c.commandLevel = "EXIT"
		default:
			flag = false
			fmt.Println("Введите 1, 2 или 3")
		}
		if flag {
			break
		}
	}
}

func (c *Command) EnterLengthCommand() {
	fmt.Println("Введите длину слова от 3 до 8:")
	var value1 string
	for {
		fmt.Scanln(&value1)
		value2, err := strconv.Atoi(value1)
		if err == nil && value2 >= 3 && value2 <= 8 {
			c.commandLength = value2
			break
		} else {
			fmt.Println("Неверно ! Введите длину слова от 3 до 8:")
		}
	}
}

func (c *Command) ReadCommandLength() int {
	return c.commandLength
}

func (c *Command) ReadCommandLevel() string {
	return c.commandLevel
}

func (c *Command) ReadCommandGame() string {
	return c.commandLevel
}
