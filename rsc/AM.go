package Hangman

import (
	"fmt"
	"math/rand"
	"os"
)

func ChoixMot() string {
	words := GetValueFromFile()
	if len(words) == 0 {
		fmt.Println("No words found.")
		return ""
	}

	randomWord := GetRandomWord(words)
	return randomWord
}

func GetValueFromFile() []string {
	content, err := os.ReadFile("Dictionnaire.txt")
	if err != nil {
		fmt.Println("Erreur de lecture")
		return []string{}
	}

	var word string
	var tab []string

	for _, c := range string(content) {
		if c >= 'A' && c <= 'Z' {
			word += string(c)
		} else if word != "" {
			tab = append(tab, word)
			word = ""
		}
	}

	if word != "" {
		tab = append(tab, word)
	}
	return tab
}

func GetRandomWord(words []string) string {
	if len(words) == 0 {
		return ""
	}

	randomIndex := rand.Intn(len(words))
	return words[randomIndex]
}
