package main

import (
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
		case 1:
		case 2:
			os.Exit(0)
		default:
			fmt.Println("Choix Invalide")
			continue
		}
	}
}