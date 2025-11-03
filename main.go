package main

import (
	"fmt"
)

func main() {
	afficherMenu()
}

func afficherMenu() {
	for {
		fmt.Println("\n=== Mini-CRM ===")
		fmt.Println("1. Ajouter un contact")
		fmt.Println("2. Lister tous les contacts")
		fmt.Println("3. Supprimer un contact")
		fmt.Println("4. Mettre à jour un contact")
		fmt.Println("5. Quitter")
		fmt.Print("\nChoisissez une option: ")

		var choix string
		fmt.Scanln(&choix)

		switch choix {
		case "1":
			fmt.Println("Ajout d'un contact...")
		case "2":
			fmt.Println("Liste des contacts...")
		case "3":
			fmt.Println("Suppression d'un contact...")
		case "4":
			fmt.Println("Mise à jour d'un contact...")
		case "5":
			fmt.Println("Au revoir!")
			return
		default:
			fmt.Println("Option invalide, veuillez réessayer.")
		}
	}
}
