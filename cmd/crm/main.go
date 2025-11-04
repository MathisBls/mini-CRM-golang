package main

import (
	"flag"
	"fmt"
	"tp1/internal/handler"
	"tp1/internal/store"
)

func main() {
	storer := store.New()
	h := handler.New(storer)
	afficherMenu(h)
}

func afficherMenu(h *handler.Handler) {
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
			h.AjouterContact()
		case "2":
			h.ListerContacts()
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
