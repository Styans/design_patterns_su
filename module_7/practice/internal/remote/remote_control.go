package remote

import (
	"fmt"
	"smart-patterns-go/internal/command"
)

type RemoteControl struct {
	onCommands  []command.ICommand
	offCommands []command.ICommand
	undoCommand command.ICommand
}

func NewRemoteControl(slots int) *RemoteControl {
	on := make([]command.ICommand, slots)
	off := make([]command.ICommand, slots)

	noCommand := &command.NoCommand{}
	for i := 0; i < slots; i++ {
		on[i] = noCommand
		off[i] = noCommand
	}

	return &RemoteControl{
		onCommands:  on,
		offCommands: off,
		undoCommand: noCommand,
	}
}

func (r *RemoteControl) SetCommand(slot int, onCmd command.ICommand, offCmd command.ICommand) {
	if slot >= 0 && slot < len(r.onCommands) {
		r.onCommands[slot] = onCmd
		r.offCommands[slot] = offCmd
	}
}

func (r *RemoteControl) PressOnButton(slot int) {
	if slot >= 0 && slot < len(r.onCommands) {
		cmd := r.onCommands[slot]
		fmt.Printf("\n[Slot %d] Нажата кнопка ON. ", slot)
		cmd.Execute()
		r.undoCommand = cmd 
	} else {
		fmt.Println("Ошибка: Неверный номер слота.")
	}
}

func (r *RemoteControl) PressOffButton(slot int) {
	if slot >= 0 && slot < len(r.offCommands) {
		cmd := r.offCommands[slot]
		fmt.Printf("\n[Slot %d] Нажата кнопка OFF. ", slot)
		cmd.Execute()
		r.undoCommand = cmd 
	} else {
		fmt.Println("Ошибка: Неверный номер слота.")
	}
}

func (r *RemoteControl) PressUndoButton() {
	fmt.Println("\n<- Нажата кнопка UNDO...")
	r.undoCommand.Undo()
}
