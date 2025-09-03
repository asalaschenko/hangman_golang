package loader

import (
	"errors"
	"log/slog"
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
	return &Loader{
		filePath: filePath}
}

func (l *Loader) SelectRandom(length int) string {
	const op = "loader.SelectRandom"
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError})).With(
		slog.String("op", op),
	)

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	arrayWord, err := l.SelectRequiredLength(length)

	if err != nil {
		log.Error("array is empty")
		return ""
	} else {
		return arrayWord[rnd.Intn(len(arrayWord))]
	}
}

func (l *Loader) SelectRequiredLength(length int) ([]string, error) {
	var arrayWord []string
	for _, v := range l.list {
		if utf8.RuneCountInString(v) == length {
			arrayWord = append(arrayWord, v)
		}
	}

	if len(arrayWord) == 0 {
		return nil, errors.New("array is empty")
	} else {
		return arrayWord, nil
	}

}
func (l *Loader) LoadWord() {
	data, err := os.ReadFile(l.filePath)
	const op = "loader.LoadWord"
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError})).With(
		slog.String("op", op),
	)

	if err != nil {
		log.Error("error of read file")
	}

	l.list = strings.Split(string(data), "\n")
}
