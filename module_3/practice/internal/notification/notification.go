package notification

type Notifier interface {
	SendNotification(message string) error
	Name() string
}

type EmailNotification struct {
	To string
}

func (e EmailNotification) SendNotification(message string) error {
	return nil
}

func (EmailNotification) Name() string {
	return "Email"
}

type SmsNotification struct {
	Phone string
}

func (s SmsNotification) SendNotification(message string) error {
	return nil
}

func (SmsNotification) Name() string {
	return "SMS"
}
