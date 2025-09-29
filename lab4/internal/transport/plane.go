package transport

import "fmt"

type Plane struct {
	Model string
	Speed int
}

func (p Plane) Move(i int) {
	fmt.Println("The Plane is moving")
}
func (p Plane) FuelUp() {
	fmt.Println("The Plane is fueling up")
}
