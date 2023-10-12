package Hangman

import "fmt"

func Menu() {
	fmt.Println("Que souhaitez vous faire ?\n 1.Pendu classic ")
	for {
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:

			Pendu("exemple")
		default:
			fmt.Println("Choix invalide, Seul 1 est accepté")
			continue
		}
	}
}
