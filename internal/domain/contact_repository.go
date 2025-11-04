package domain

type Storer interface {
	Ajouter(contact *Contact)
	Lister() map[int]*Contact
	ProchainID() int
}
