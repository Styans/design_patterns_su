package command

import "fmt"

type ICommand interface {
	Execute()
	Undo()
	Redo()
}

type NoCommand struct{}

func (c *NoCommand) Execute() {
	fmt.Println("🚫 Слот пуст. Ничего не назначено.")
}
func (c *NoCommand) Undo() {}
func (c *NoCommand) Redo() {}