package mediator

import "fmt"

type Colleague interface {
	GetName() string
	ReceiveMessage(message string, senderName string)
}

type User struct {
	Name     string
	Mediator IMediator
}

func NewUser(name string, mediator IMediator) *User {
	user := &User{Name: name, Mediator: mediator}
	mediator.Register(user)
	return user
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) Send(message string) {
	fmt.Printf("\n%s отправляет сообщение: %s\n", u.Name, message)
	u.Mediator.SendMessage(message, u)
}

func (u *User) ReceiveMessage(message string, senderName string) {
	fmt.Printf("  [%s получил от %s]: %s\n", u.Name, senderName, message)
}