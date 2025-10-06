package ocp

type Employee struct {
	Name       string
	BaseSalary float64
}

type SalaryCalculator interface {
	CalculateSalary() float64
}

type PermanentEmployee struct {
	Employee
}

func (p PermanentEmployee) CalculateSalary() float64 {
	return p.BaseSalary * 1.2
}

type ContractEmployee struct {
	Employee
}

func (c ContractEmployee) CalculateSalary() float64 {
	return c.BaseSalary * 1.1
}

type Intern struct {
	Employee
}

func (i Intern) CalculateSalary() float64 {
	return i.BaseSalary * 0.8
}
