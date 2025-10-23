package main

import (
	"fmt"
	"home/internal/command"
	"home/internal/mediator"
	"home/internal/templatemethod"
)

func main() {
	fmt.Println("===================================================")
	fmt.Println("            ДЕМОНСТРАЦИЯ ПАТТЕРНОВ GO              ")
	fmt.Println("===================================================")

	fmt.Println("\n--- 1. ПАТТЕРН КОМАНДА: УМНЫЙ ДОМ (С UNDO) ---")
	demoCommand()

	fmt.Println("\n--- 2. ПАТТЕРН ШАБЛОННЫЙ МЕТОД: НАПИТКИ ---")
	demoTemplateMethod()

	fmt.Println("\n--- 3. ПАТТЕРН ПОСРЕДНИК: ЧАТ-КОМНАТА ---")
	demoMediator()

	fmt.Println("\n===================================================")
}

func demoCommand() {
	light := command.NewLight()
	door := command.NewDoor()
	thermostat := command.NewThermostat(20)

	invoker := command.NewSmartHomeInvoker()
	invoker.CurrentState(light, door, thermostat)

	invoker.ExecuteCommand(command.NewLightOnCommand(light))
	invoker.ExecuteCommand(command.NewTempIncreaseCommand(thermostat))
	invoker.ExecuteCommand(command.NewDoorOpenCommand(door))
	invoker.CurrentState(light, door, thermostat)

	invoker.UndoLastCommand()
	invoker.CurrentState(light, door, thermostat)

	invoker.UndoLastCommand()
	invoker.CurrentState(light, door, thermostat)

	invoker.UndoLastCommand()
	invoker.UndoLastCommand()
	invoker.UndoLastCommand()
	invoker.CurrentState(light, door, thermostat)
}

func demoTemplateMethod() {
	fmt.Println("--- Готовим Кофе: ---")
	coffee := templatemethod.NewCoffee()
	coffee.MakeBeverage() 

	fmt.Println("--- Готовим Чай: ---")
	tea := templatemethod.NewTea()
	tea.MakeBeverage() 
}

func demoMediator() {
	chatroom := mediator.NewChatRoom()

	userA := mediator.NewUser("Алиса", chatroom)
	userB := mediator.NewUser("Боб", chatroom)
	userC := mediator.NewUser("Чарли", chatroom)

	userA.Send("Привет, всем!")
	userB.Send("Рад вас видеть, Чарли, Алиса.")

	userC.Send("Всем привет! У меня все отлично, спасибо.")

	userD := &mediator.User{Name: "Дэвид", Mediator: chatroom}
	userD.Send("Я пытаюсь отправить сообщение, но не зарегистрирован.")
}
