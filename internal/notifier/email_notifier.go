package notifier

import "fmt"

type EmailNotifier struct {
	From string
}

func (e EmailNotifier) Send(message string) error {
	fmt.Printf("[EMAIL] From: %s | Message: %s\n", e.From, message)
	return nil
}
