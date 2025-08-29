package core

type IDisplay interface {
	DisplayMenu()
	DisplayHelp()
}

type IGetCommand interface {
	EnterLevelCommand()
	EnterGameCommand()
	EnterLengthCommand()
	ReadCommandLevel() string
	ReadCommandGame() string
	ReadCommandLength() int
}

type IGameLoop interface {
	StartGameLoop()
	SetWord(string)
	SetCountOfAttempts(string)
	DisplayResult()
	GetResultGame() string
}

type ILoader interface {
	SelectRandom(int) string
	LoaderWord()
}
