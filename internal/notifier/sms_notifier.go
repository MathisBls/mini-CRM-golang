package notifier

import "fmt"

type SmsNotifier struct {
	Number string
}

func (s SmsNotifier) Send(message string) error {
	fmt.Printf("[SMS] To: %s | Message: %s\n", s.Number, message)
	return nil
}
