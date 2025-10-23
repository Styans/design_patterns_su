package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"smart-patterns-go/internal/command"
	"smart-patterns-go/internal/mediator"
	"smart-patterns-go/internal/templatemethod"
)

func demoCommand() {
	fmt.Println("=====================================")
	fmt.Println("1. ПАТТЕРН КОМАНДА: УМНЫЙ ДОМ")
	fmt.Println("=====================================")

	light := command.NewLight()
	tv := command.NewTelevision()
	thermostat := command.NewThermostat()

	// Объявленные команды теперь используются
	lightOn := command.NewLightOnCommand(light)
	lightOff := command.NewLightOffCommand(light)
	tvOn := command.NewTelevisionOnCommand(tv)
	tvOff := command.NewTelevisionOffCommand(tv)
	tempUp := command.NewTempUpCommand(thermostat, 2)
	tempDown := command.NewTempDownCommand(thermostat, 2)

	eveningMacro := command.NewMacroCommand([]command.ICommand{lightOn, tvOn, tempUp})

	remote := command.NewRemoteControl()

	fmt.Println("\n--- Управление светом ---")
	remote.SetCommand(lightOn)
	remote.PressButton() // Включаем
	remote.SetCommand(lightOff)
	remote.PressButton() // Выключаем
	remote.Undo()        // Отменяем выключение (должен включиться)

	fmt.Println("\n--- Управление TV ---")
	remote.SetCommand(tvOn)
	remote.PressButton() // Включаем
	remote.SetCommand(tvOff)
	remote.PressButton() // Выключаем
	remote.Undo()        // Отменяем выключение (должен включиться)

	fmt.Println("\n--- Управление термостатом ---")
	remote.SetCommand(tempUp)
	remote.PressButton() // +2
	remote.SetCommand(tempDown)
	remote.PressButton() // -2
	remote.Undo()        // Отменяем -2 (должен быть +2)

	fmt.Println("\n--- Запуск Макроса 'Вечер' (Свет On, TV On, Темп +2) ---")
	remote.SetCommand(eveningMacro)
	remote.PressButton()
	remote.Undo() // Отмена всех 3 команд макроса
}

func demoTemplateMethod() {
	fmt.Println("\n=====================================")
	fmt.Println("2. ПАТТЕРН ШАБЛОННЫЙ МЕТОД: НАПИТКИ")
	fmt.Println("=====================================")

	reader := bufio.NewReader(os.Stdin)

	tea := templatemethod.NewTea()
	fmt.Println("\n--- Приготовление Чая ---")
	tea.PrepareRecipe()

	coffee := templatemethod.NewCoffee(func() bool {
		fmt.Print("  Хотите добавить сахар и молоко (да/нет)? ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		return input == "да" || input == "yes"
	})
	fmt.Println("\n--- Приготовление Кофе ---")
	coffee.PrepareRecipe()
}

func demoMediator() {
	fmt.Println("\n=====================================")
	fmt.Println("3. ПАТТЕРН ПОСРЕДНИК: ЧАТ-СИСТЕМА")
	fmt.Println("=====================================")

	chatroom := mediator.NewChatRoom()

	userA := mediator.NewUser("Алиса", chatroom)
	userB := mediator.NewUser("Боб", chatroom)
	userC := mediator.NewUser("Чарли", chatroom)

	userA.Send("Привет, всем!")
	userB.Send("Рад вас видеть.")
	userC.Send("Всем привет! У меня все отлично, спасибо.")

	userD := &mediator.User{Name: "Дэвид (не зарег.)", Mediator: chatroom}
	userD.Send("Я пытаюсь отправить сообщение, но не зарегистрирован.")
}

func main() {
	demoCommand()
	demoTemplateMethod()
	demoMediator()
}