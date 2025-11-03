package main

import "testing"

func TestAjouterContact(t *testing.T) {
	contacts = make(map[int]Contact)
	prochainID = 1

	contact := Contact{ID: 1, Nom: "Jean", Email: "jean@test.com"}
	contacts[1] = contact

	if len(contacts) != 1 {
		t.Errorf("len(contacts) = %d; want 1", len(contacts))
	}

	_, existe := contacts[1]
	if !existe {
		t.Error("Contact avec ID 1 devrait exister")
	}
}

func TestListerContacts(t *testing.T) {
	contacts = make(map[int]Contact)

	contacts[1] = Contact{ID: 1, Nom: "Alice", Email: "alice@test.com"}
	contacts[2] = Contact{ID: 2, Nom: "Bob", Email: "bob@test.com"}

	if len(contacts) != 2 {
		t.Errorf("len(contacts) = %d; want 2", len(contacts))
	}
}
