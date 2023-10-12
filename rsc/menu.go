package Hangman

import (
	"fmt"
	"os"
)

func Menu() {
	ClearTerminal()
	for {
		fmt.Println("Que souhaitez vous faire ?\n1.Pendu avec tous les mot du dictionaire\n2.Facile\n3.Moyen\n4.Difficile\n0.Partir\n ")

		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 0:
			os.Exit(0)
		case 1:
			s := ChoixMot(choice)
			Pendu(s)
		case 2:
			s := ChoixMot(choice)
			Pendu(s)
		case 3:
			s := ChoixMot(choice)
			Pendu(s)
		case 4:
			s := ChoixMot(choice)
			Pendu(s)

		default:
			fmt.Println("Choix invalide, Seul 1 est accepté")
			continue
		}
	}
}
