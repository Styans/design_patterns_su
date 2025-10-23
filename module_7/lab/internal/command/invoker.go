package command

import "fmt"

type NoCommand struct{}

func (c *NoCommand) Execute() {
	fmt.Println("[ERROR]: Кнопка не настроена.")
}

func (c *NoCommand) Undo() {
	fmt.Println("[ERROR]: Нечего отменять.")
}

type RemoteControl struct {
	command ICommand
	history ICommand
}

func NewRemoteControl() *RemoteControl {
	return &RemoteControl{history: &NoCommand{}}
}

func (r *RemoteControl) SetCommand(cmd ICommand) {
	r.command = cmd
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
	r.history = r.command
}

func (r *RemoteControl) Undo() {
	fmt.Println("  [UNDO]: Отмена последней команды...")
	r.history.Undo()
	r.history = &NoCommand{}
}