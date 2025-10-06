package vehicles

type Fleet struct {
	Garages []Garage
}

func (f *Fleet) AddGarage(g Garage) {
	f.Garages = append(f.Garages, g)
}

func (f *Fleet) RemoveGarage(index int) {
	if index >= 0 && index < len(f.Garages) {
		f.Garages = append(f.Garages[:index], f.Garages[index+1:]...)
	}
}

func (f *Fleet) FindVehicle(brand, model string) Vehicle {
	for _, garage := range f.Garages {
		for _, v := range garage.Vehicles {
			if v.GetBrand() == brand && v.GetModel() == model {
				return v
			}
		}
	}
	return nil
}
