package store

import (
	"testing"
	"tp1/internal/domain"
)

func TestAjouterContact(t *testing.T) {
	s := New()
	s.Reset()

	contact := &domain.Contact{ID: 1, Nom: "Jean", Email: "jean@test.com"}
	s.Ajouter(contact)

	contacts := s.Lister()
	if len(contacts) != 1 {
		t.Errorf("len(contacts) = %d; want 1", len(contacts))
	}

	_, existe := contacts[1]
	if !existe {
		t.Error("Contact avec ID 1 devrait exister")
	}
}

func TestListerContacts(t *testing.T) {
	s := New()
	s.Reset()

	s.Ajouter(&domain.Contact{ID: 1, Nom: "Alice", Email: "alice@test.com"})
	s.Ajouter(&domain.Contact{ID: 2, Nom: "Bob", Email: "bob@test.com"})

	contacts := s.Lister()
	if len(contacts) != 2 {
		t.Errorf("len(contacts) = %d; want 2", len(contacts))
	}
}
