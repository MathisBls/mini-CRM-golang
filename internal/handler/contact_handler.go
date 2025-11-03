package handler

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func SupprimerContact() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("ID du contact à supprimer : ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Erreur : l'ID doit être un nombre.")
		return
	}

	ok := store.Supprimer(id)
	if ok {
		fmt.Println("Contact supprimé avec succès.")
	} else {
		fmt.Println("Aucun contact trouvé avec cet ID.")
	}
}

func ModifierContact() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("ID du contact à modifier : ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Erreur : l'ID doit être un nombre.")
		return
	}

	fmt.Print("Nouveau nom (laisser vide pour ne pas changer) : ")
	nom, _ := reader.ReadString('\n')
	nom = strings.TrimSpace(nom)

	fmt.Print("Nouvel email (laisser vide pour ne pas changer) : ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	ok := store.Modifier(id, nom, email)
	if ok {
		fmt.Println("Contact modifié avec succès.")
	} else {
		fmt.Println("Aucun contact trouvé avec cet ID.")
	}
}

func AjouterContactViaFlags(id int, nom string, email string) {
	contact := domain.Contact{
		ID:    id,
		Nom:   nom,
		Email: email,
	}
	store.Ajouter(contact)
	fmt.Printf("Contact ajouté via flags : ID=%d, Nom=%s, Email=%s\n", id, nom, email)
}
