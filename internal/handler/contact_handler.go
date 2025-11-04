package handler

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func (h *Handler) SupprimerContact() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("ID du contact à supprimer : ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Erreur : l'ID doit être un nombre.")
		return
	}

	ok := h.storer.Supprimer(id)
	if ok {
		fmt.Println("Contact supprimé avec succès.")
	} else {
		fmt.Println("Aucun contact trouvé avec cet ID.")
	}
}

func (h *Handler) ModifierContact() {
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

	ok := h.storer.Modifier(id, nom, email)
	if ok {
		fmt.Println("Contact modifié avec succès.")
	} else {
		fmt.Println("Aucun contact trouvé avec cet ID.")
	}
}
