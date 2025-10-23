package mediator

import "fmt"

type Mediator interface {
	SendMessage(message string, sender *User)
	Register(user *User)
}

type ChatRoom struct {
	users map[string]*User
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{users: make(map[string]*User)}
}

func (c *ChatRoom) Register(user *User) {
	if _, exists := c.users[user.Name]; exists {
		fmt.Printf("   [ChatRoom]: Пользователь %s уже зарегистрирован.\n", user.Name)
		return
	}
	c.users[user.Name] = user
	fmt.Printf("   [ChatRoom]: Пользователь %s присоединился к чату.\n", user.Name)
	c.SendSystemNotification(fmt.Sprintf("%s присоединился к чату.", user.Name), user.Name)
}

func (c *ChatRoom) SendMessage(message string, sender *User) {
	if _, exists := c.users[sender.Name]; !exists {
		fmt.Printf("   [ChatRoom ERROR]: Пользователь %s не в чате. Отправка невозможна.\n", sender.Name)
		return
	}

	fmt.Printf("\n[ОБЩИЙ ЧАТ] %s: %s\n", sender.Name, message)
	for name, user := range c.users {
		if name != sender.Name {
			user.Receive(sender.Name, message)
		}
	}
}

func (c *ChatRoom) SendSystemNotification(message string, excludedUser string) {
	for name, user := range c.users {
		if name != excludedUser {
			user.Receive("СИСТЕМА", message)
		}
	}
}