package main

import (
	Hangman "Hangman/rsc"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Bienvenue sur notre jeu de Pendu")
	fmt.Println("Que voulez vous faire ? (1.jouer / 2. Partir)")
	for {
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 0:
			Hangman.GetValueFromFile()
		case 1:
			Hangman.Menu()
		case 2:
			os.Exit(0)
		default:
			fmt.Println("Choix Invalide")
			continue
		}
	}
}
