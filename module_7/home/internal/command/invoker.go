package command

import "fmt"

type SmartHomeInvoker struct {
	history []Command
}

func NewSmartHomeInvoker() *SmartHomeInvoker {
	return &SmartHomeInvoker{history: make([]Command, 0)}
}

func (i *SmartHomeInvoker) ExecuteCommand(cmd Command) {
	fmt.Println("  [Invoker] Выполнение команды...")
	cmd.Execute()
	i.history = append(i.history, cmd)
}

func (i *SmartHomeInvoker) UndoLastCommand() {
	if len(i.history) == 0 {
		fmt.Println("  [Invoker] Ошибка: История команд пуста, нечего отменять.")
		return
	}
	lastCommand := i.history[len(i.history)-1]
	i.history = i.history[:len(i.history)-1]

	fmt.Println("  [Invoker] Отмена последней команды...")
	lastCommand.Undo()
}

func (i *SmartHomeInvoker) CurrentState(l *Light, d *Door, t *Thermostat) {
	fmt.Printf("  [STATUS]: Свет=%v, Дверь=%v, Темп=%d°C, Команд в истории: %d\n",
		l.IsOn, d.IsOpen, t.Temperature, len(i.history))
}