package Hangman

import "fmt"

func Menu() {
	fmt.Println("Souhaitez vous faire 1.Pendu classic ?")
	for {
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
		default:
			fmt.Println("Choix invalide, Seul 1 ou 2 sont accepté")
			continue
		}
	}
}
