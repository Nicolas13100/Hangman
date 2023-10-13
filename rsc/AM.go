package Hangman

import (
	"fmt"
	"math/rand"
	"os"
)

func ChoixMot(choix int) string {
	words := GetValueFromFile(choix)
	if len(words) == 0 {
		fmt.Println("No words found.")
		return ""
	}
	randomWord := GetRandomWord(words)
	return randomWord
}

func GetValueFromFile(choix int) []string {
	switch choix {
	case 1:
		content, err := os.ReadFile("Librairie/Dictionnaire.txt")
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
	case 2:
		content, err := os.ReadFile("Librairie/Facile.txt")
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
	case 3:
		content, err := os.ReadFile("Librairie/Moyen.txt")
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
	case 4:
		content, err := os.ReadFile("Librairie/Difficile.txt")
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
	case 5:
		content, err := os.ReadFile("Librairie/Halloween.txt")
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
	return nil
}

func GetRandomWord(words []string) string {
	if len(words) == 0 {
		return ""
	}

	randomIndex := rand.Intn(len(words))
	return words[randomIndex]
}
