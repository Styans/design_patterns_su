package vehicles

import "fmt"

type Car struct {
	Brand       string
	Model       string
	Year        int
	Doors       int
	Transmission string
}

func (c Car) GetBrand() string { return c.Brand }
func (c Car) GetModel() string { return c.Model }
func (c Car) GetYear() int     { return c.Year }

func (c Car) StartEngine() string {
	return fmt.Sprintf("%s %s engine started", c.Brand, c.Model)
}

func (c Car) StopEngine() string {
	return fmt.Sprintf("%s %s engine stopped", c.Brand, c.Model)
}
