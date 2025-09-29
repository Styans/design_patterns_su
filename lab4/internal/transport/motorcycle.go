package transport

import "fmt"

type Motorcycle struct {
	Model string
	Speed int
}

func (m Motorcycle) Move(i int) {
	fmt.Println("The motorcycle is moving")
}

func (m Motorcycle) FuelUp() {
	fmt.Println("The motorcycle is fueling up")

}
