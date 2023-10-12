package Hangman

import (
	"fmt"
	"os"
	"strings"
)

var hangman []string

func init() {
	// Read hangman ASCII art from file
	data, err := os.ReadFile("hangman.txt")
	if err != nil {
		fmt.Println("Error lecture hangman file:", err)
		return
	}
	hangman = strings.Split(string(data), "\n")
}

func printHangman(incorrectGuesses int) {
	if incorrectGuesses < len(hangman) {
		fmt.Println(hangman[incorrectGuesses])
	} else {
		fmt.Println(hangman[len(hangman)-1])
	}
}

func Pendu(s string) {
	fmt.Println("C'est parti, a vous de jouer")
	word := s
	guessedLetters := []string{}
	incorrectGuesses := 0

	for {
		var user_input string
		fmt.Print("Entrer un lettre: ")
		_, err := fmt.Scan(&user_input)

		if err != nil || len(user_input) != 1 || !strings.ContainsAny(user_input, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
			fmt.Println("Entrer une lettre valide SVP.")
			continue
		}

		result, guessedLetters := checkLetter(word, guessedLetters, strings.ToLower(user_input))

		if !result {
			incorrectGuesses++
		}

		printHangman(incorrectGuesses)

		displayWord := ""
		for _, letter := range word {
			if contains(guessedLetters, string(letter)) {
				displayWord += string(letter) + " "
			} else {
				displayWord += "_ "
			}
		}

		fmt.Println(displayWord)

		fmt.Printf("Lettre donné: %s\n", strings.Join(guessedLetters, ", "))

		if strings.Contains(word, strings.Join(guessedLetters, "")) {
			fmt.Printf("Bravo, vous avez trouver le mot: %s\n", word)
			break
		}

		if incorrectGuesses == 6 {
			fmt.Printf("Dommage, plus de tentative. Le mot été: %s\n", word)
			break
		}
	}
}

func contains(slice []string, item string) bool {
	for _, element := range slice {
		if element == item {
			return true
		}
	}
	return false
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
