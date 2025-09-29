package employee

type Manager struct {
	Name        string
	ID          int
	Position    string
	FixedSalary float64
	Bonus       float64
}

func (m Manager) GetName() string {
	return m.Name
}

func (m Manager) GetPosition() string {
	return m.Position
}

func (m Manager) CalculateSalary() float64 {
	return m.FixedSalary + m.Bonus
}
