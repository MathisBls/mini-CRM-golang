package store

import "tp1/internal/domain"

var contacts map[int]domain.Contact = make(map[int]domain.Contact)
var prochainID int = 1

func Ajouter(contact domain.Contact) {
	contacts[contact.ID] = contact
}

func Lister() map[int]domain.Contact {
	return contacts
}

func ProchainID() int {
	id := prochainID
	prochainID++
	return id
}

func Reset() {
	contacts = make(map[int]domain.Contact)
	prochainID = 1
}

func Supprimer(id int) bool {
	_, existe := contacts[id]
	if existe {
		delete(contacts, id)
		return true
	}
	return false
}

func Modifier(id int, nom string, email string) bool {
	contact, existe := contacts[id]
	if !existe {
		return false
	}

	if nom != "" {
		contact.Nom = nom
	}
	if email != "" {
		contact.Email = email
	}
	contacts[id] = contact
	return true
}

