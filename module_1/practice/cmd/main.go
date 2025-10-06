package main

import (
	"fmt"
	vehicles "pr/internal/vehicle"
)

func main() {
	car1 := vehicles.Car{Brand: "Toyota", Model: "Camry", Year: 2020, Doors: 4, Transmission: "Automatic"}
	car2 := vehicles.Car{Brand: "BMW", Model: "X5", Year: 2021, Doors: 5, Transmission: "Manual"}
	bike1 := vehicles.Motorcycle{Brand: "Yamaha", Model: "R1", Year: 2019, Body: "Sport", HasBox: false}

	garage1 := vehicles.Garage{Name: "Garage A"}
	garage1.AddVehicle(car1)
	garage1.AddVehicle(bike1)

	garage2 := vehicles.Garage{Name: "Garage B"}
	garage2.AddVehicle(car2)

	fleet := vehicles.Fleet{}
	fleet.AddGarage(garage1)
	fleet.AddGarage(garage2)

	fmt.Println(car1.StartEngine())
	fmt.Println(bike1.StartEngine())

	found := fleet.FindVehicle("BMW", "X5")
	if found != nil {
		fmt.Printf("Found: %s %s\n", found.GetBrand(), found.GetModel())
	}

	garage1.RemoveVehicle(0)
	fleet.RemoveGarage(1)
}
