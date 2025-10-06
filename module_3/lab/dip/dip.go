package dip

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailService struct{}

func (EmailService) Send(message string) {
	fmt.Println("Отправка Email:", message)
}

type SmsService struct{}

func (SmsService) Send(message string) {
	fmt.Println("Отправка SMS:", message)
}

type Notification struct {
	notifier Notifier
}

func NewNotification(notifier Notifier) Notification {
	return Notification{notifier: notifier}
}

func (n Notification) Send(message string) {
	n.notifier.Send(message)
}
