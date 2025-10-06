package vehicle

import "fmt"

type IVehicle interface {
	Drive() string
	Refuel() string
}

type Car struct {
	Brand string
	Model string
	Fuel  string
}

func (c Car) Drive() string {
	return "Car " + c.Brand + " " + c.Model + " is driving"
}

func (c Car) Refuel() string {
	return "Car " + c.Brand + " refueled with " + c.Fuel
}

type Motorcycle struct {
	Type   string
	Engine int
}

func (m Motorcycle) Drive() string {
	return "Motorcycle " + m.Type + " is driving"
}

func (m Motorcycle) Refuel() string {
	return "Motorcycle refueled"
}

type Truck struct {
	LoadCapacity float64
	Axles        int
}

func (t Truck) Drive() string {
	return "Truck with capacity " + fmt.Sprintf("%.1f", t.LoadCapacity) + " tons is driving"
}

func (t Truck) Refuel() string {
	return "Truck refueled"
}

type Bus struct {
	Seats int
	Route string
}

func (b Bus) Drive() string {
	return "Bus on route " + b.Route + " is driving with " + fmt.Sprintf("%d", b.Seats) + " seats"
}

func (b Bus) Refuel() string {
	return "Bus refueled"
}

type Scooter struct {
	Model string
	Range int
}

func (s Scooter) Drive() string {
	return "Scooter " + s.Model + " is driving with range " + fmt.Sprintf("%d", s.Range) + " km"
}

func (s Scooter) Refuel() string {
	return "Scooter charged"
}
