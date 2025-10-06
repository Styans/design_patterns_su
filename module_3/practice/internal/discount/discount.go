package discount

type DiscountStrategy interface {
	Apply(amount float64) float64
}

type NoDiscount struct{}

func (NoDiscount) Apply(amount float64) float64 {
	return amount
}

type PercentageDiscount struct {
	Percent float64
}

func (p PercentageDiscount) Apply(amount float64) float64 {
	return amount * (1 - p.Percent/100)
}

type DiscountCalculator struct {
	strategies []DiscountStrategy
}

func NewCalculator(strats ...DiscountStrategy) DiscountCalculator {
	return DiscountCalculator{strategies: strats}
}

func (c DiscountCalculator) Calculate(amount float64) float64 {
	res := amount
	for _, s := range c.strategies {
		res = s.Apply(res)
	}
	return res
}
