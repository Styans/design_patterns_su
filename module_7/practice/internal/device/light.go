package device

import "fmt"

type Light struct {
	Name  string
	State bool 
}

func NewLight(name string) *Light {
	return &Light{Name: name, State: false}
}

func (l *Light) On() {
	l.State = true
	fmt.Printf("%s: 💡 Свет включен.\n", l.Name)
}

func (l *Light) Off() {
	l.State = false
	fmt.Printf("%s: 💡 Свет выключен.\n", l.Name)
}