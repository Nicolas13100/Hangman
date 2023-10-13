package Hangman

import (
	"fmt"
	"os"
	"strings"
)

var hangman []string

func Init() {
	// Read hangman ASCII art from file
	data, err := os.ReadFile("rsc/hangman.txt")
	if err != nil {
		fmt.Println("Error reading hangman file:", err)
		return
	}

	if len(data) == 0 {
		fmt.Println("Error: hangman file is empty")
		return
	}
	hangman = strings.Split(string(data), "\n")
}

func PrintHangman(incorrectGuesses int) {
	if incorrectGuesses == 0 {
		fmt.Println(hangman[0])
	} else if incorrectGuesses <= len(hangman) {
		start := (incorrectGuesses - 1) * 7
		end := incorrectGuesses * 7

		if end > len(hangman) {
			end = len(hangman)
		}

		for i := start; i < end; i++ {
			fmt.Println(hangman[i])
		}
	} else {
		fmt.Println(hangman[len(hangman)-1])
	}
}

func Pendu(s string) {
	fmt.Println("C'est parti, a vous de jouer")
	word := s
	GuessedLetters := []string{}
	incorrectGuesses := 0
	result := false

	for {
		var user_input string
		if len(GuessedLetters) > 0 {
			fmt.Printf("Lettre deja donné: %s\n", strings.Join(GuessedLetters, ", "))
		}
		fmt.Print("Entrer un lettre ou un mot: ")
		_, err := fmt.Scan(&user_input)

		if err != nil {
			fmt.Println("Erreur de saisie:", err)
			continue
		}

		if len(user_input) == 1 && strings.ContainsAny(user_input, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
			result, GuessedLetters = checkLetter(word, GuessedLetters, strings.ToUpper(user_input))

			if !result {
				incorrectGuesses++
			}
		} else if len(user_input) == len(word) && strings.ContainsAny(user_input, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
			USER := strings.ToUpper(user_input)
			if USER == word {
				fmt.Printf("Bravo, vous avez trouvé le mot: %s\n", word)
				break
			} else {
				incorrectGuesses += 2
			}
		} else {
			fmt.Println("Entrer une lettre ou un mot valide SVP.")
			continue
		}

		PrintHangman(incorrectGuesses)

		displayWord := ""
		for _, letter := range word {
			if contains(GuessedLetters, string(letter)) {
				displayWord += string(letter) + " "
			} else {
				displayWord += "_ "
			}
		}

		fmt.Println(displayWord)

		if strings.Join(strings.Fields(displayWord), "") == word {
			fmt.Printf("Bravo, vous avez trouvé le mot: %s\n", word)
			break
		}

		if incorrectGuesses == 9 {
			fmt.Printf("Dommage, plus de tentative. Le mot était: %s\n", word)
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

func checkLetter(word string, GuessedLetters []string, letter string) (bool, []string) {
	for _, guessedLetter := range GuessedLetters {
		if guessedLetter == letter {
			fmt.Printf("Vous avez déja essayer la lettre '%s'.\n", letter)
			return false, GuessedLetters
		}
	}

	if strings.Contains(word, letter) {
		GuessedLetters = append(GuessedLetters, letter)
		return true, GuessedLetters
	} else {
		GuessedLetters = append(GuessedLetters, letter)
		return false, GuessedLetters
	}
}
