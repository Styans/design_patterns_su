package main

import (
	"lab/internal/employee"
)

func main() {
	system := employee.EmployeeSystem{}

	worker1 := employee.Worker{Name: "Иван", ID: 1, Position: "Рабочий", HourlyRate: 500, Hours: 160}
	worker2 := employee.Worker{Name: "Петр", ID: 2, Position: "Рабочий", HourlyRate: 450, Hours: 170}
	manager1 := employee.Manager{Name: "Анна", ID: 3, Position: "Менеджер", FixedSalary: 60000, Bonus: 15000}

	system.AddEmployee(worker1)
	system.AddEmployee(worker2)
	system.AddEmployee(manager1)

	system.ShowSalaries()
}
