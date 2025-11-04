package domain

type Storer interface {
	Ajouter(contact *Contact)
	Lister() map[int]*Contact
	ProchainID() int
	Reset()
	Supprimer(id int) bool
	Modifier(id int, nom string, email string) bool
}
