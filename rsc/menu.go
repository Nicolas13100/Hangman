package Hangman

import (
	"fmt"
	"os"
)

func Menu() {
	for {
		fmt.Println("Que souhaitez vous faire ?\n 1.Pendu classic 0.Partir ")

		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 0:
			os.Exit(0)
		case 1:
			s := ChoixMot()
			Pendu(s)
		default:
			fmt.Println("Choix invalide, Seul 1 est accepté")
			continue
		}
	}
}
