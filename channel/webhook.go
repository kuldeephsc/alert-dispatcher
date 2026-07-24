package channel

import (
	"fmt"

	"github.com/kuldeephsc/model"
)

type WebhookChannel struct{}

func (w *WebhookChannel) Send(alert model.Alert) error {
	// Simulate sending a webhook
	fmt.Printf("Sending Webhook: %s - %s\n", alert.Title, alert.Message)
	return nil
}

func (w *WebhookChannel) Name() string {
	return "Webhook"
}
