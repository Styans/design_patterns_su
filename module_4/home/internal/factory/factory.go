package factory

import (
	"home/internal/vehicle"
)

type VehicleFactory interface {
	CreateVehicle() vehicle.IVehicle
}

type CarFactory struct {
	Brand string
	Model string
	Fuel  string
}

func (f CarFactory) CreateVehicle() vehicle.IVehicle {
	return vehicle.Car{Brand: f.Brand, Model: f.Model, Fuel: f.Fuel}
}

type MotorcycleFactory struct {
	Type   string
	Engine int
}

func (f MotorcycleFactory) CreateVehicle() vehicle.IVehicle {
	return vehicle.Motorcycle{Type: f.Type, Engine: f.Engine}
}

type TruckFactory struct {
	LoadCapacity float64
	Axles        int
}

func (f TruckFactory) CreateVehicle() vehicle.IVehicle {
	return vehicle.Truck{LoadCapacity: f.LoadCapacity, Axles: f.Axles}
}

type BusFactory struct {
	Seats int
	Route string
}

func (f BusFactory) CreateVehicle() vehicle.IVehicle {
	return vehicle.Bus{Seats: f.Seats, Route: f.Route}
}

type ScooterFactory struct {
	Model string
	Range int
}

func (f ScooterFactory) CreateVehicle() vehicle.IVehicle {
	return vehicle.Scooter{Model: f.Model, Range: f.Range}
}
