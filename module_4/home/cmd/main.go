package main

import (
	"fmt"
	"home/internal/factory"
)

func main() {
	var choice int
	fmt.Println("Выберите тип транспорта:")
	fmt.Println("1 - Автомобиль")
	fmt.Println("2 - Мотоцикл")
	fmt.Println("3 - Грузовик")
	fmt.Println("4 - Автобус")
	fmt.Println("5 - Электросамокат")
	fmt.Print("Ваш выбор: ")
	fmt.Scan(&choice)

	var v interface{}
	switch choice {
	case 1:
		var brand, model, fuel string
		fmt.Print("Введите марку: ")
		fmt.Scan(&brand)
		fmt.Print("Введите модель: ")
		fmt.Scan(&model)
		fmt.Print("Введите тип топлива: ")
		fmt.Scan(&fuel)
		f := factory.CarFactory{Brand: brand, Model: model, Fuel: fuel}
		v = f.CreateVehicle()
	case 2:
		var vtype string
		var engine int
		fmt.Print("Введите тип мотоцикла (спортивный/туристический): ")
		fmt.Scan(&vtype)
		fmt.Print("Введите объем двигателя: ")
		fmt.Scan(&engine)
		f := factory.MotorcycleFactory{Type: vtype, Engine: engine}
		v = f.CreateVehicle()
	case 3:
		var capacity float64
		var axles int
		fmt.Print("Введите грузоподъемность (тонны): ")
		fmt.Scan(&capacity)
		fmt.Print("Введите количество осей: ")
		fmt.Scan(&axles)
		f := factory.TruckFactory{LoadCapacity: capacity, Axles: axles}
		v = f.CreateVehicle()
	case 4:
		var seats int
		var route string
		fmt.Print("Введите количество мест: ")
		fmt.Scan(&seats)
		fmt.Print("Введите маршрут: ")
		fmt.Scan(&route)
		f := factory.BusFactory{Seats: seats, Route: route}
		v = f.CreateVehicle()
	case 5:
		var model string
		var rng int
		fmt.Print("Введите модель самоката: ")
		fmt.Scan(&model)
		fmt.Print("Введите запас хода (км): ")
		fmt.Scan(&rng)
		f := factory.ScooterFactory{Model: model, Range: rng}
		v = f.CreateVehicle()
	default:
		fmt.Println("Неверный выбор")
		return
	}

	if veh, ok := v.(interface {
		Drive() string
		Refuel() string
	}); ok {
		fmt.Println(veh.Drive())
		fmt.Println(veh.Refuel())
	}
}
