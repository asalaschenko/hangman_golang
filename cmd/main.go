package main

func main() {
	game := CreateDI()
	game.StartGame()
}

// func CreateDI() *core.Core {
// 	loaderWordObj := loader.NewLoader("./resource/words.txt")

// 	gameLoopObj := game.NewGame()

// 	displayObj := display.NewDisplay("./resource/help.txt", "Приветствую в игре Виселица !\n", "ГЛАВНОЕ МЕНЮ")

// 	commandObj := command.NewCommand()

// 	coreObj := core.NewCore(displayObj, commandObj, gameLoopObj, loaderWordObj)

// 	return coreObj
// }
