package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Contact struct {
	ID    int
	Nom   string
	Email string
}

var contacts map[int]Contact = make(map[int]Contact)
var prochainID int = 1

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
			ajouterContact()
		case "2":
			listerContacts()
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

func ajouterContact() {
	reader := bufio.NewReader(os.Stdin)

	id := prochainID
	prochainID++

	fmt.Print("Nom: ")
	nom, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erreur lors de la lecture du nom:", err)
		return
	}
	nom = strings.TrimSpace(nom)

	fmt.Print("Email: ")
	email, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erreur lors de la lecture de l'email:", err)
		return
	}
	email = strings.TrimSpace(email)

	contact := Contact{
		ID:    id,
		Nom:   nom,
		Email: email,
	}

	contacts[id] = contact
	fmt.Printf("Contact ajouté avec succès! (ID: %d)\n", id)
}

func listerContacts() {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact enregistré.")
		return
	}

	fmt.Println("\n=== Liste des contacts ===")
	for id, contact := range contacts {
		fmt.Printf("ID: %d | Nom: %s | Email: %s\n", id, contact.Nom, contact.Email)
	}
}
