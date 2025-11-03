package beverage

// 1. Базовый интерфейс Beverage
type Beverage interface {
	GetDescription() string
	Cost() float64
}

// --- 2. Конкретные напитки ---

type Espresso struct{}

func (e *Espresso) GetDescription() string { return "Эспрессо" }
func (e *Espresso) Cost() float64          { return 1.99 }

type Tea struct{}

func (t *Tea) GetDescription() string { return "Чай" }
func (t *Tea) Cost() float64          { return 1.50 }

type Latte struct{}

func (l *Latte) GetDescription() string { return "Латте" }
func (l *Latte) Cost() float64          { return 2.49 }

type Mocha struct{}

func (m *Mocha) GetDescription() string { return "Мокка" }
func (m *Mocha) Cost() float64          { return 2.79 }

// --- 3 & 4. Декораторы (реализованы через композицию) ---

// Milk
type Milk struct {
	beverage Beverage // ссылка на оборачиваемый объект (неэкспортируемая)
}

func NewMilk(b Beverage) *Milk {
	return &Milk{beverage: b}
}

func (m *Milk) GetDescription() string {
	return m.beverage.GetDescription() + ", Молоко"
}
func (m *Milk) Cost() float64 {
	return m.beverage.Cost() + 0.50
}

// Sugar
type Sugar struct {
	beverage Beverage
}

func NewSugar(b Beverage) *Sugar {
	return &Sugar{beverage: b}
}

func (s *Sugar) GetDescription() string {
	return s.beverage.GetDescription() + ", Сахар"
}
func (s *Sugar) Cost() float64 {
	return s.beverage.Cost() + 0.20
}

// WhippedCream
type WhippedCream struct {
	beverage Beverage
}

func NewWhippedCream(b Beverage) *WhippedCream {
	return &WhippedCream{beverage: b}
}

func (w *WhippedCream) GetDescription() string {
	return w.beverage.GetDescription() + ", Взбитые сливки"
}
func (w *WhippedCream) Cost() float64 {
	return w.beverage.Cost() + 0.70
}

// Syrup (Добавка)
type Syrup struct {
	beverage Beverage
}

func NewSyrup(b Beverage) *Syrup {
	return &Syrup{beverage: b}
}

func (s *Syrup) GetDescription() string {
	return s.beverage.GetDescription() + ", Сироп"
}
func (s *Syrup) Cost() float64 {
	return s.beverage.Cost() + 0.60
}

// Cinnamon (Добавка)
type Cinnamon struct {
	beverage Beverage
}

func NewCinnamon(b Beverage) *Cinnamon {
	return &Cinnamon{beverage: b}
}

func (c *Cinnamon) GetDescription() string {
	return c.beverage.GetDescription() + ", Корица"
}
func (c *Cinnamon) Cost() float64 {
	return c.beverage.Cost() + 0.30
}