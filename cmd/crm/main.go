package main

import (
	"flag"
	"fmt"
	"tp1/internal/handler"
)

func main() {
	addFlag := flag.Bool("add", false, "Ajouter un contact via des flags")
	idFlag := flag.Int("id", 0, "ID du contact")
	nomFlag := flag.String("nom", "", "Nom du contact")
	emailFlag := flag.String("email", "", "Email du contact")

	flag.Parse()

	if *addFlag {
		if *idFlag == 0 || *nomFlag == "" || *emailFlag == "" {
			fmt.Println("Utilisation : go run cmd/crm/main.go -add -id=1 -nom=\"Kevin\" -email=\"kevin@exemple.com\"")
			return
		}
		handler.AjouterContactViaFlags(*idFlag, *nomFlag, *emailFlag)
		return
	}

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
			handler.AjouterContact()
		case "2":
			handler.ListerContacts()
		case "3":
			handler.SupprimerContact()
		case "4":
			handler.ModifierContact()
		case "5":
			fmt.Println("Au revoir!")
			return
		default:
			fmt.Println("Option invalide, veuillez réessayer.")
		}
	}
}
