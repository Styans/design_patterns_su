package templatemethod

import "fmt"

type BrewHook func() bool

type Beverage interface {
	PrepareRecipe()
	boilWater()
	pourInCup()
	brew()
	addCondiments()
	customerWantsCondiments() bool
}

type baseBeverage struct {
	Beverage
}

func (b *baseBeverage) PrepareRecipe() {
	b.boilWater()
	b.brew()
	b.pourInCup()
	if b.customerWantsCondiments() {
		b.addCondiments()
	}
}

func (b *baseBeverage) boilWater() {
	fmt.Println("1. Кипячение воды.")
}

func (b *baseBeverage) pourInCup() {
	fmt.Println("3. Наливание в чашку.")
}