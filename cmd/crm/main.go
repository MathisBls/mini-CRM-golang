package main

import (
	"flag"
	"fmt"
	"tp1/internal/handler"
	"tp1/internal/store"
	"tp1/internal/notifier"
)

func main() {
	flag.Parse()

	storer := store.New()
	h := handler.New(storer)

	emailNotifier := notifier.EmailNotifier{From: "crm@entreprise.com"}
	smsNotifier := notifier.SmsNotifier{Number: "+33612345678"}
	notifiers := []notifier.Notifier{emailNotifier, smsNotifier}

	afficherMenu(h, notifiers)
}

func afficherMenu(h *handler.Handler, notifiers []notifier.Notifier) {
	for {
		fmt.Println("\n=== Mini-CRM ===")
		fmt.Println("1. Ajouter un contact")
		fmt.Println("2. Lister tous les contacts")
		fmt.Println("3. Supprimer un contact")
		fmt.Println("4. Mettre à jour un contact")
		fmt.Println("5. Quitter")
		fmt.Println("6. Envoyer une notification")
		fmt.Print("\nChoisissez une option: ")

		var choix string
		fmt.Scanln(&choix)

		switch choix {
		case "1":
			h.AjouterContact()
		case "2":
			h.ListerContacts()
		case "3":
			h.SupprimerContact()
		case "4":
			h.ModifierContact()
		case "5":
			fmt.Println("Au revoir!")
			return
		case "6":
			envoyerNotification(notifiers)
		default:
			fmt.Println("Option invalide, veuillez réessayer.")
		}
	}
}

func envoyerNotification(notifiers []notifier.Notifier) {
	var message string
	fmt.Print("Entrez le message à envoyer : ")
	fmt.Scanln(&message)

	for _, n := range notifiers {
		if err := n.Send(message); err != nil {
			fmt.Println("Erreur lors de l'envoi :", err)
		}
	}
}
