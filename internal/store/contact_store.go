package store

import "tp1/internal/domain"

type Store struct {
	contacts   map[int]*domain.Contact
	prochainID int
}

func New() *Store {
	return &Store{
		contacts:   make(map[int]*domain.Contact),
		prochainID: 1,
	}
}

func (s *Store) Ajouter(contact *domain.Contact) {
	s.contacts[contact.ID] = contact
}

func (s *Store) Lister() map[int]*domain.Contact {
	return s.contacts
}

func (s *Store) ProchainID() int {
	id := s.prochainID
	s.prochainID++
	return id
}

func (s *Store) Reset() {
	s.contacts = make(map[int]*domain.Contact)
	s.prochainID = 1
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

