package Hangman

import (
	"fmt"
	"strings"
)

func Pendu(s string) {
	fmt.Println("C'est parti, a vous de jouer")
	word := s
	guessedLetters := []string{}

	for {
		var user_input string
		fmt.Print("Entrez une lettre: ")
		_, err := fmt.Scan(&user_input)

		if err != nil || len(user_input) != 1 {
			fmt.Println("N'entrez qu'une lettre a la fois SVP.")
			continue
		}

		result, guessedLetters := checkLetter(word, guessedLetters, user_input)

		if result {
			fmt.Printf("Bien joué! '%s' est dans le mot.\n", user_input)
		} else {
			fmt.Printf("Désolé, '%s' n'est pas dans le mot.\n", user_input)
		}

		fmt.Printf("Lettre donné: %s\n", strings.Join(guessedLetters, ", "))

		if strings.Contains(word, strings.Join(guessedLetters, "")) {
			fmt.Printf("Bravo, vous avez trouvé le mot: %s\n", word)
			break
		}
	}
}

func checkLetter(word string, guessedLetters []string, letter string) (bool, []string) {
	for _, guessedLetter := range guessedLetters {
		if guessedLetter == letter {
			fmt.Printf("Vous avez déja essayer la lettre '%s'.\n", letter)
			return false, guessedLetters
		}
	}

	if strings.Contains(word, letter) {
		guessedLetters = append(guessedLetters, letter)
		return true, guessedLetters
	} else {
		guessedLetters = append(guessedLetters, letter)
		return false, guessedLetters
	}
}
