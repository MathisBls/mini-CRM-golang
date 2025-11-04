package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tp1/internal/domain"
)

type Handler struct {
	storer domain.Storer
}

func New(storer domain.Storer) *Handler {
	return &Handler{storer: storer}
}

func (h *Handler) AjouterContact() {
	reader := bufio.NewReader(os.Stdin)

	id := h.storer.ProchainID()

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

	contact := &domain.Contact{
		ID:    id,
		Nom:   nom,
		Email: email,
	}

	h.storer.Ajouter(contact)
	fmt.Printf("Contact ajouté avec succès! (ID: %d)\n", id)
}

func (h *Handler) ListerContacts() {
	contacts := h.storer.Lister()

	if len(contacts) == 0 {
		fmt.Println("Aucun contact enregistré.")
		return
	}

	fmt.Println("\n=== Liste des contacts ===")
	for id, contact := range contacts {
		fmt.Printf("ID: %d | Nom: %s | Email: %s\n", id, contact.Nom, contact.Email)
	}
}
