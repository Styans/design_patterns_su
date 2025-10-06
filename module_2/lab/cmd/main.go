package main

import (
	"lab/internal/dry"
	"lab/internal/kiss"
	"lab/internal/yagni"
	"fmt"
)

func main() {
	order := dry.OrderService{}
	order.CreateOrder("Book", 3, 12.5)
	order.UpdateOrder("Book", 5, 12.5)

	car := dry.Car{Vehicle: dry.Vehicle{Type: "Car"}}
	truck := dry.Truck{Vehicle: dry.Vehicle{Type: "Truck"}}
	car.Start()
	car.Stop()
	truck.Start()
	truck.Stop()

	calculator := kiss.Calculator{}
	calculator.Add(5, 10)

	client := kiss.Client{}
	client.Execute()

	circle := yagni.Circle{Radius: 4.2}
	fmt.Printf("Circle area: %.2f\n", circle.CalculateArea())

	math := yagni.MathOperations{}
	fmt.Println("Sum:", math.Add(7, 3))
}
