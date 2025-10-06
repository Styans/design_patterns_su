package dip

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailSender struct{}

func (e EmailSender) Send(message string) {
	fmt.Println("Email sent:", message)
}

type SmsSender struct{}

func (s SmsSender) Send(message string) {
	fmt.Println("SMS sent:", message)
}

type NotificationService struct {
	notifiers []Notifier
}

func NewNotificationService(notifiers ...Notifier) NotificationService {
	return NotificationService{notifiers: notifiers}
}

func (n NotificationService) SendNotification(message string) {
	for _, notifier := range n.notifiers {
		notifier.Send(message)
	}
}
