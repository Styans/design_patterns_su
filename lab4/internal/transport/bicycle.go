package transport

import "fmt"

type Bicycle struct {
	Model string
	Speed int
}

func (p Bicycle) Move(i int) {
	fmt.Println("The Bicycle is moving")
}
func (p Bicycle) FuelUp() {
	fmt.Println("The Bicycle is fueling up")
}
