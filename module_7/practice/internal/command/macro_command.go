package command

import "fmt"

type MacroCommand struct {
	commands []ICommand
}

func NewMacroCommand(cmds []ICommand) *MacroCommand {
	return &MacroCommand{commands: cmds}
}

func (c *MacroCommand) Execute() {
	fmt.Println("\n--- Выполнение Макрокоманды ---")
	for _, cmd := range c.commands {
		cmd.Execute()
	}
	fmt.Println("--- Макрокоманда завершена ---")
}

func (c *MacroCommand) Undo() {
	fmt.Println("\n--- Отмена Макрокоманды ---")
	for i := len(c.commands) - 1; i >= 0; i-- {
		c.commands[i].Undo()
	}
	fmt.Println("--- Отмена завершена ---")
}

func (c *MacroCommand) Redo() {
	c.Execute()
}