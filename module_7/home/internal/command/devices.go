package command

import "fmt"

type Light struct {
	IsOn bool
}

func NewLight() *Light { return &Light{} }

func (l *Light) TurnOn() {
	l.IsOn = true
	fmt.Println("   [Light]: Свет включен.")
}

func (l *Light) TurnOff() {
	l.IsOn = false
	fmt.Println("   [Light]: Свет выключен.")
}

type Door struct {
	IsOpen bool
}

func NewDoor() *Door { return &Door{} }

func (d *Door) Open() {
	d.IsOpen = true
	fmt.Println("   [Door]: Дверь открыта.")
}

func (d *Door) Close() {
	d.IsOpen = false
	fmt.Println("   [Door]: Дверь закрыта.")
}

type Thermostat struct {
	Temperature int
}

func NewThermostat(initialTemp int) *Thermostat {
	return &Thermostat{Temperature: initialTemp}
}

func (t *Thermostat) IncreaseTemp() {
	t.Temperature++
	fmt.Printf("   [Thermostat]: Температура увеличена до %d°C.\n", t.Temperature)
}

func (t *Thermostat) DecreaseTemp() {
	t.Temperature--
	fmt.Printf("   [Thermostat]: Температура уменьшена до %d°C.\n", t.Temperature)
}