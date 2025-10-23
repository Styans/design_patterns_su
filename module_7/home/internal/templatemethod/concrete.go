package templatemethod

import (
	"fmt"
	"strings"
)

type Coffee struct {
	BaseBeverage
}

func NewCoffee() *Coffee {
	c := &Coffee{}
	c.BaseBeverage = NewBaseBeverage(c)
	return c
}

func (c *Coffee) Brew() {
	fmt.Println("2. Заваривание кофе через фильтр")
}

func (c *Coffee) AddCondiments() {
	fmt.Println("4. Добавление сахара и молока")
}

func (c *Coffee) CustomerWantsCondiments() bool {
	fmt.Print("   Хотите добавить сахар и молоко (да/нет)? ")
	var answer string
	fmt.Scanln(&answer)
	if strings.ToLower(answer) == "да" || strings.ToLower(answer) == "yes" {
		return true
	}
	return false
}

type Tea struct {
	BaseBeverage
}

func NewTea() *Tea {
	t := &Tea{}
	t.BaseBeverage = NewBaseBeverage(t)
	return t
}

func (t *Tea) Brew() {
	fmt.Println("2. Замачивание чайного пакетика")
}

func (t *Tea) AddCondiments() {
	fmt.Println("4. Добавление лимона")
}
