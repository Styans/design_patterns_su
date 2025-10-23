package templatemethod

import "fmt"

type Tea struct {
	baseBeverage
}

func NewTea() *Tea {
	tea := &Tea{}
	tea.baseBeverage.Beverage = tea
	return tea
}

func (t *Tea) brew() {
	fmt.Println("2. Заваривание чая.")
}

func (t *Tea) addCondiments() {
	fmt.Println("4. Добавление лимона.")
}

func (t *Tea) customerWantsCondiments() bool {
	return true
}

type Coffee struct {
	baseBeverage
	hook BrewHook
}

func NewCoffee(hook BrewHook) *Coffee {
	coffee := &Coffee{hook: hook}
	coffee.baseBeverage.Beverage = coffee
	return coffee
}

func (c *Coffee) brew() {
	fmt.Println("2. Заваривание кофе.")
}

func (c *Coffee) addCondiments() {
	fmt.Println("4. Добавление сахара и молока.")
}

func (c *Coffee) customerWantsCondiments() bool {
	if c.hook != nil {
		return c.hook()
	}
	return true
}