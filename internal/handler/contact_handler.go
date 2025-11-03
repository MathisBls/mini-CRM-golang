package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tp1/internal/domain"
	"tp1/internal/store"
)

func AjouterContact() {
	reader := bufio.NewReader(os.Stdin)

	id := store.ProchainID()

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

	contact := domain.Contact{
		ID:    id,
		Nom:   nom,
		Email: email,
	}

	store.Ajouter(contact)
	fmt.Printf("Contact ajouté avec succès! (ID: %d)\n", id)
}

func ListerContacts() {
	contacts := store.Lister()

	if len(contacts) == 0 {
		fmt.Println("Aucun contact enregistré.")
		return
	}

	fmt.Println("\n=== Liste des contacts ===")
	for id, contact := range contacts {
		fmt.Printf("ID: %d | Nom: %s | Email: %s\n", id, contact.Nom, contact.Email)
	}
}
