package employee

type Worker struct {
	Name       string
	ID         int
	Position   string
	HourlyRate float64
	Hours      float64
}

func (w Worker) GetName() string {
	return w.Name
}

func (w Worker) GetPosition() string {
	return w.Position
}

func (w Worker) CalculateSalary() float64 {
	return w.HourlyRate * w.Hours
}
