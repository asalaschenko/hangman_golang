package loader

import (
	"hangman/common"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

type Loader struct {
	filePath string
	list     []string
}

func NewLoader(filePath string) *Loader {
	value1 := &Loader{
		filePath: filePath}
	return value1
}

func (l *Loader) SelectRandom(length int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		value := l.list[rnd.Intn(len(l.list))]
		if utf8.RuneCountInString(value) == length {
			return value
		}
	}
}

func (l *Loader) LoaderWord() {
	data, err := os.ReadFile(l.filePath)
	common.CheckErr(err)

	list := string(data)
	l.list = strings.Split(list, "\n")
}
