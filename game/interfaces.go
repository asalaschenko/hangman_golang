package game

type Display interface {
	DisplayMenu()
	DisplayHelp()
	DisplayStartPhrase()
}

type Loader interface {
	SelectRandom(int) string
	LoadWord()
}

type GetCommand interface {
	EnterLevelCommand()
	EnterGameCommand()
	EnterLengthCommand()
	ReadCommandLevel() string
	ReadCommandGame() string
	ReadCommandLength() int
}
