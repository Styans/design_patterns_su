package practice

import "fmt"

// OrganizationComponent — общий интерфейс для сотрудников и отделов.
type OrganizationComponent interface {
	GetName() string
	Display(indent int)
	GetBudget() float64
	GetHeadcount() int
}

// вспомогательная функция для отступов
func indentStr(indent int) string {
	s := ""
	for i := 0; i < indent; i++ {
		s += "-"
	}
	return s
}

// -------------------- Сотрудник --------------------

type Employee struct {
	name     string
	position string
	salary   float64
}

func NewEmployee(name, position string, salary float64) *Employee {
	return &Employee{name: name, position: position, salary: salary}
}

func (e *Employee) GetName() string {
	return e.name
}

func (e *Employee) Display(indent int) {
	fmt.Printf("%s Employee: %s (%s), salary=%.2f\n",
		indentStr(indent), e.name, e.position, e.salary)
}

func (e *Employee) GetBudget() float64 {
	return e.salary
}

func (e *Employee) GetHeadcount() int {
	return 1
}

func (e *Employee) SetSalary(salary float64) {
	e.salary = salary
}

// -------------------- Контрактор (не входит в бюджет отдела) --------------------

type Contractor struct {
	name     string
	role     string
	fee      float64
}

func NewContractor(name, role string, fee float64) *Contractor {
	return &Contractor{name: name, role: role, fee: fee}
}

func (c *Contractor) GetName() string {
	return c.name
}

func (c *Contractor) Display(indent int) {
	fmt.Printf("%s Contractor: %s (%s), fee=%.2f [NOT in dept budget]\n",
		indentStr(indent), c.name, c.role, c.fee)
}

// Зарплата контракторов НЕ входит в бюджет отдела
func (c *Contractor) GetBudget() float64 {
	return 0.0
}

func (c *Contractor) GetHeadcount() int {
	return 1
}

// -------------------- Отдел (Department) --------------------

type Department struct {
	name     string
	children []OrganizationComponent
}

func NewDepartment(name string) *Department {
	return &Department{name: name, children: make([]OrganizationComponent, 0)}
}

func (d *Department) GetName() string {
	return d.name
}

func (d *Department) Add(component OrganizationComponent) {
	d.children = append(d.children, component)
}

func (d *Department) Remove(component OrganizationComponent) {
	for i, c := range d.children {
		if c == component {
			d.children = append(d.children[:i], d.children[i+1:]...)
			break
		}
	}
}

func (d *Department) Display(indent int) {
	fmt.Printf("%s Department: %s\n", indentStr(indent), d.name)
	for _, c := range d.children {
		c.Display(indent + 2)
	}
}

func (d *Department) GetBudget() float64 {
	total := 0.0
	for _, c := range d.children {
		total += c.GetBudget()
	}
	return total
}

func (d *Department) GetHeadcount() int {
	total := 0
	for _, c := range d.children {
		total += c.GetHeadcount()
	}
	return total
}

// Поиск сотрудника/контрактора по имени (возвращаем первый найденный компонент)
func (d *Department) FindByName(name string) OrganizationComponent {
	if d.name == name {
		return d
	}
	for _, c := range d.children {
		if c.GetName() == name {
			return c
		}
		// Рекурсивный поиск в подотделах
		if subDept, ok := c.(*Department); ok {
			if found := subDept.FindByName(name); found != nil {
				return found
			}
		}
	}
	return nil
}

// Вывод всех сотрудников (и контракторов) отдела и подотделов
func (d *Department) ListAllEmployees() {
	fmt.Printf("=== Employees of department %s and subdepartments ===\n", d.name)
	d.listEmployeesRecursive(0)
}

func (d *Department) listEmployeesRecursive(indent int) {
	for _, c := range d.children {
		switch v := c.(type) {
		case *Employee, *Contractor:
			c.Display(indent + 2)
		case *Department:
			v.listEmployeesRecursive(indent + 2)
		}
	}
}
