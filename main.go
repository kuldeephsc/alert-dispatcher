package main

import (
	"github.com/kuldeephsc/channel"
	"github.com/kuldeephsc/dispatcher"
	"github.com/kuldeephsc/service"
)

func main() {
	d := dispatcher.NewDispatcher()
	d.Register(&channel.EmailChannel{})
	d.Register(&channel.SlackChannel{})
	d.Register(&channel.WebhookChannel{})
	notification := service.NewNotificationService(d)
	notification.Notify("Test Alert", "This is a test notification message.")

}
