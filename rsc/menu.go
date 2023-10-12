package Hangman

import "fmt"

func Menu() {
	fmt.Println("Souhaitez vous faire 1.Pendu d'alloween ou 2.Pendu classic ?")
	for {
		var choice int
		switch choice {
		case 1:
			PenduAlloween()
		case 2:
			Pendu()
		default:
			fmt.Println("Choix invalide, Seul 1 ou 2 sont accepté")
			continue
		}
	}
}
