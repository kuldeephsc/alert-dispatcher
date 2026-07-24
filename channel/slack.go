package channel

import (
	"fmt"

	"github.com/kuldeephsc/model"
)

type SlackChannel struct{}

func (s *SlackChannel) Send(alert model.Alert) error {
	// Simulate sending a Slack message
	fmt.Printf("Sending Slack: %s - %s\n", alert.Title, alert.Message)
	return nil
}

func (s *SlackChannel) Name() string {
	return "Slack"
}
