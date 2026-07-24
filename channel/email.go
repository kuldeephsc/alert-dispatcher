package channel

import (
	"fmt"

	"github.com/kuldeephsc/model"
)

type EmailChannel struct{}

func (e *EmailChannel) Send(alert model.Alert) error {
	// Simulate sending an email
	fmt.Printf("Sending Email: %s - %s\n", alert.Title, alert.Message)
	return nil
}

func (e *EmailChannel) Name() string {
	return "Email"
}
