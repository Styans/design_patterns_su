package command

import "fmt"

type Light struct{}

func NewLight() *Light { return &Light{} }

func (l *Light) On() {
	fmt.Println("[Light]: Свет включен.")
}

func (l *Light) Off() {
	fmt.Println("[Light]: Свет выключен.")
}

type Television struct{}

func NewTelevision() *Television { return &Television{} }

func (t *Television) On() {
	fmt.Println("[TV]: Телевизор включен.")
}

func (t *Television) Off() {
	fmt.Println("[TV]: Телевизор выключен.")
}

type Thermostat struct {
	temperature int
}

func NewThermostat() *Thermostat {
	return &Thermostat{temperature: 20}
}

func (t *Thermostat) Up(degrees int) {
	t.temperature += degrees
	fmt.Printf("[Thermostat]: Температура поднята. Текущая: %d°C\n", t.temperature)
}

func (t *Thermostat) Down(degrees int) {
	t.temperature -= degrees
	fmt.Printf("[Thermostat]: Температура опущена. Текущая: %d°C\n", t.temperature)
}

func (t *Thermostat) GetTemp() int {
	return t.temperature
}