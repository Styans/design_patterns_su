package templatemethod

import "fmt"

type Beverage interface {
	Brew()
	AddCondiments()
	CustomerWantsCondiments() bool 
	MakeBeverage()               
}

type BaseBeverage struct {
	Beverage
}

func NewBaseBeverage(b Beverage) BaseBeverage {
	return BaseBeverage{Beverage: b}
}

func (b *BaseBeverage) MakeBeverage() {
	b.BoilWater()
	b.Beverage.Brew()
	b.PourInCup()
	if b.Beverage.CustomerWantsCondiments() { 
		b.Beverage.AddCondiments()
	}
	fmt.Println("--- Напиток готов! ---\n")
}

func (b *BaseBeverage) BoilWater() {
	fmt.Println("1. Кипячение воды")
}

func (b *BaseBeverage) PourInCup() {
	fmt.Println("3. Наливаем в чашку")
}

func (b *BaseBeverage) CustomerWantsCondiments() bool {
	return true
}