package ocp

type DiscountStrategy interface {
	ApplyDiscount(amount float64) float64
}

type RegularCustomer struct{}

func (RegularCustomer) ApplyDiscount(amount float64) float64 {
	return amount
}

type SilverCustomer struct{}

func (SilverCustomer) ApplyDiscount(amount float64) float64 {
	return amount * 0.9
}

type GoldCustomer struct{}

func (GoldCustomer) ApplyDiscount(amount float64) float64 {
	return amount * 0.8
}

type DiscountCalculator struct {
	strategy DiscountStrategy
}

func NewDiscountCalculator(strategy DiscountStrategy) DiscountCalculator {
	return DiscountCalculator{strategy: strategy}
}

func (dc DiscountCalculator) Calculate(amount float64) float64 {
	return dc.strategy.ApplyDiscount(amount)
}
