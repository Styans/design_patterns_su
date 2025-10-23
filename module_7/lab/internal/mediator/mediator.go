package mediator

import (
	"fmt"
	"strings"
)

type IMediator interface {
	SendMessage(message string, sender Colleague)
	Register(colleague Colleague)
}

type ChatRoom struct {
	colleagues []Colleague
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{}
}

func (c *ChatRoom) Register(colleague Colleague) {
	for _, regColleague := range c.colleagues {
		if strings.EqualFold(regColleague.GetName(), colleague.GetName()) {
			fmt.Printf("[СИСТЕМА]: Участник %s уже зарегистрирован.\n", colleague.GetName())
			return
		}
	}
	c.colleagues = append(c.colleagues, colleague)
	fmt.Printf("[СИСТЕМА]: Участник %s присоединился.\n", colleague.GetName())
}

func (c *ChatRoom) SendMessage(message string, sender Colleague) {
	isRegistered := false
	for _, regColleague := range c.colleagues {
		if regColleague == sender {
			isRegistered = true
			break
		}
	}

	if !isRegistered {
		fmt.Printf("[ОШИБКА]: Участник %s не зарегистрирован и не может отправлять сообщения.\n", sender.GetName())
		return
	}

	for _, colleague := range c.colleagues {
		if colleague != sender {
			colleague.ReceiveMessage(message, sender.GetName())
		}
	}
}