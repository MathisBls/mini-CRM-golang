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
