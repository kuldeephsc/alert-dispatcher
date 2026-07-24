package service

import (
	"fmt"

	"github.com/kuldeephsc/dispatcher"
	"github.com/kuldeephsc/model"
)

type NotificationService struct {
	dispatcher *dispatcher.Dispatcher
}

func NewNotificationService(dispatcher *dispatcher.Dispatcher) *NotificationService {
	return &NotificationService{dispatcher: dispatcher}
}

func (ns *NotificationService) Notify(title, message string) {
	alert := model.Alert{
		Title:   title,
		Message: message,
	}
	result := ns.dispatcher.Dispatch(alert)
	fmt.Println("Notification Results:", result)
}
