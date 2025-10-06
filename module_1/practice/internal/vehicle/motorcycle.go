package vehicles

import "fmt"

type Motorcycle struct {
	Brand   string
	Model   string
	Year    int
	Body    string
	HasBox  bool
}

func (m Motorcycle) GetBrand() string { return m.Brand }
func (m Motorcycle) GetModel() string { return m.Model }
func (m Motorcycle) GetYear() int     { return m.Year }

func (m Motorcycle) StartEngine() string {
	return fmt.Sprintf("%s %s engine started", m.Brand, m.Model)
}

func (m Motorcycle) StopEngine() string {
	return fmt.Sprintf("%s %s engine stopped", m.Brand, m.Model)
}
