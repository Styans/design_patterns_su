package main

import (
	"fmt"
	"lab4/internal/factory"
	"strconv"
	"strings"
)

func main() {
	var choice string
	fmt.Println("Выберите тип транспорта: car / motorcycle / plane / bicycle")
	fmt.Scan(&choice)

	choice = strings.ToLower(choice)

	switch choice {
	case "car", "motorcycle", "plane", "bicycle":
	default:
		fmt.Println("Ошибка: неизвестный тип транспорта")
		return
	}
	
	var model, fuel string
	var speed string

	fmt.Print("Введите модель: ")
	fmt.Scan(&model)

	fmt.Print("Введите скорость (км/ч): ")
	fmt.Scan(&speed)

	_, err := strconv.Atoi(speed)
	if err != nil {
		fmt.Println("Ошибка: скорость должна быть числом")
		return
	}

	if choice != "bicycle" {
		fmt.Print("Введите топливо: ")
		fmt.Scan(&fuel)
		_, err := strconv.Atoi(fuel)
		if err != nil {
			fmt.Println("Ошибка: скорость должна быть числом")
			return
		}
	}
	f := factory.GetFactory(choice)

	t := f.CreateTransport()
	t.Move(10)
	t.FuelUp()
}
