package main

import (
	"fmt"
	"smart-patterns-go/internal/command"
	"smart-patterns-go/internal/device"
	"smart-patterns-go/internal/remote"
)

func main() {
	fmt.Println("=== 🏠 Система Умный Дом (Паттерн Command) ===")

	remoteControl := remote.NewRemoteControl(4)

	kitchenLight := device.NewLight("Кухня")
	livingRoomLight := device.NewLight("Гостиная")
	bedroomAC := device.NewAirConditioner("Спальня", 25)

	remoteControl.SetCommand(0, command.NewLightOnCommand(kitchenLight), command.NewLightOffCommand(kitchenLight))

	acCool := command.NewAcSetTempCommand(bedroomAC, 20)
	acWarm := command.NewAcSetTempCommand(bedroomAC, 28)
	remoteControl.SetCommand(1, acCool, acWarm)

	fmt.Println("\n--- ТЕСТ: Выполнение и Отмена ---")

	remoteControl.PressOnButton(0) // Кухня ON
	remoteControl.PressOnButton(1) // AC на 20°C

	remoteControl.PressUndoButton() // Отмена AC 20°C -> 25°C

	remoteControl.PressOffButton(0) // Кухня OFF
	remoteControl.PressUndoButton() // Отмена Кухня OFF -> Кухня ON

	fmt.Println("\n--- ТЕСТ: Макрокоманда (Cinema Mode) ---")

	cinemaCommands := []command.ICommand{
		command.NewLightOffCommand(livingRoomLight),
		command.NewLightOnCommand(kitchenLight),
		command.NewAcSetTempCommand(bedroomAC, 22),
	}
	cinemaMacro := command.NewMacroCommand(cinemaCommands)

	// Назначаем макрос на Слот 2
	remoteControl.SetCommand(2, cinemaMacro, command.NewLightOnCommand(livingRoomLight))

	remoteControl.PressOnButton(2) // Запуск макроса

	// Отмена макроса
	remoteControl.PressUndoButton() // Отмена всех 3 команд в обратном порядке

	// Тестирование: Обработка ошибок (Пустой слот)
	fmt.Println("\n--- ТЕСТ: Пустой слот (NoCommand) ---")
	remoteControl.PressOnButton(3) // Слот 3 пуст
}
