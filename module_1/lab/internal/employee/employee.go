package employee

type Employee interface {
	GetName() string
	GetPosition() string
	CalculateSalary() float64
}
