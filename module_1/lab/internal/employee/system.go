package employee

import "fmt"

type EmployeeSystem struct {
	Employees []Employee
}

func (es *EmployeeSystem) AddEmployee(e Employee) {
	es.Employees = append(es.Employees, e)
}

func (es *EmployeeSystem) ShowSalaries() {
	for _, emp := range es.Employees {
		fmt.Printf("%s (%s) - Salary: %.2f\n", emp.GetName(), emp.GetPosition(), emp.CalculateSalary())
	}
}
