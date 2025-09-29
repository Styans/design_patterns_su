package transport

import "fmt"

type Car struct {
	Model string
	Speed int
}

func (c Car) Move(i int) {
	fmt.Println("The car is moving")
}
func (c Car) FuelUp() {
	fmt.Println("The car is fueling up")
}
