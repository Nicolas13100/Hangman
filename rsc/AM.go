package Hangman

import (
	"fmt"
	"os"
)

func ChoixMots() {

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
	fmt.Println(tab)
	return tab
}
