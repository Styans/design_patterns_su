package yagni

type Circle struct {
	Radius float64
}

func (c Circle) CalculateArea() float64 {
	return 3.14159 * c.Radius * c.Radius
}

type MathOperations struct{}

func (m MathOperations) Add(a, b int) int {
	return a + b
}
