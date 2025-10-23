package device

import "fmt"

type AirConditioner struct {
	Name        string
	Temperature int
}

func NewAirConditioner(name string, initialTemp int) *AirConditioner {
	return &AirConditioner{Name: name, Temperature: initialTemp}
}

func (ac *AirConditioner) SetTemperature(temp int) {
	fmt.Printf("%s: 🌡️ Температура изменена с %d°C на %d°C.\n", ac.Name, ac.Temperature, temp)
	ac.Temperature = temp
}

func (ac *AirConditioner) TurnOn() {
	fmt.Printf("%s: ❄️ Кондиционер включен.\n", ac.Name)
}

func (ac *AirConditioner) TurnOff() {
	fmt.Printf("%s: ❄️ Кондиционер выключен.\n", ac.Name)
}