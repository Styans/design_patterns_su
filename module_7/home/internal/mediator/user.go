package mediator

import "fmt"

type User struct {
	Name string
	Mediator Mediator
}

func NewUser(name string, m Mediator) *User {
	u := &User{Name: name, Mediator: m}
	m.Register(u) 
	return u
}

func (u *User) Send(message string) {
	u.Mediator.SendMessage(message, u)
}

func (u *User) Receive(senderName, message string) {
	fmt.Printf("  [%s -> %s]: %s\n", senderName, u.Name, message)
}