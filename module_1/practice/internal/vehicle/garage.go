package vehicles

type Garage struct {
	Name     string
	Vehicles []Vehicle
}

func (g *Garage) AddVehicle(v Vehicle) {
	g.Vehicles = append(g.Vehicles, v)
}

func (g *Garage) RemoveVehicle(index int) {
	if index >= 0 && index < len(g.Vehicles) {
		g.Vehicles = append(g.Vehicles[:index], g.Vehicles[index+1:]...)
	}
}
