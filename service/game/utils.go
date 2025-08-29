package game

import (
	"regexp"
	"strings"
)

func IsRussianLetter(char rune) bool {
	// Регулярное выражение для русских букв (включая ё)
	russianPattern := regexp.MustCompile(`^[а-яёА-ЯЁ]$`)
	return russianPattern.MatchString(string(char))
}

func IsEnglishLetter(char rune) bool {
	// Регулярное выражение для английских букв
	russianPattern := regexp.MustCompile(`^[a-zA-Z]$`)
	return russianPattern.MatchString(string(char))
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

func ReturnNumberOfAttempts(s string) int {
	switch s {
	case "HARD":
		return HARD
	case "MEDIUM":
		return MEDIUM
	case "EASY":
		return EASY
	default:
		return 0
	}
}
