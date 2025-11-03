package beverage

// 1. Базовый интерфейс IBeverage
type IBeverage interface {
	GetCost() float64
	GetDescription() string
}

// --- 2. Базовые напитки ---

// Coffee (как в примере C#)
type Coffee struct{}

func (c *Coffee) GetCost() float64 {
	return 50.0 // Стоимость кофе
}
func (c *Coffee) GetDescription() string {
	return "Coffee"
}

// Добавим еще один базовый напиток для тестов
type Tea struct{}

func (t *Tea) GetCost() float64 {
	return 40.0 // Стоимость чая
}
func (t *Tea) GetDescription() string {
	return "Tea"
}

/*
 * 3. Абстрактный декоратор:
 * В Go нет абстрактных классов. Мы используем композицию.
 * Вместо наследования от BeverageDecorator, каждый декоратор
 * будет реализовывать интерфейс IBeverage и содержать
 * вложенный объект IBeverage.
 */

// --- 4. Декораторы для добавок ---

// MilkDecorator (Молоко)
type MilkDecorator struct {
	beverage IBeverage // ссылка на "оборачиваемый" объект
}

func NewMilkDecorator(b IBeverage) *MilkDecorator {
	return &MilkDecorator{beverage: b}
}

func (m *MilkDecorator) GetCost() float64 {
	return m.beverage.GetCost() + 10.0 // + стоимость молока
}
func (m *MilkDecorator) GetDescription() string {
	return m.beverage.GetDescription() + ", Milk"
}

// SugarDecorator (Сахар)
type SugarDecorator struct {
	beverage IBeverage
}

func NewSugarDecorator(b IBeverage) *SugarDecorator {
	return &SugarDecorator{beverage: b}
}

func (s *SugarDecorator) GetCost() float64 {
	return s.beverage.GetCost() + 5.0 // + стоимость сахара
}
func (s *SugarDecorator) GetDescription() string {
	return s.beverage.GetDescription() + ", Sugar"
}

// ChocolateDecorator (Шоколад) - из нового задания
type ChocolateDecorator struct {
	beverage IBeverage
}

func NewChocolateDecorator(b IBeverage) *ChocolateDecorator {
	return &ChocolateDecorator{beverage: b}
}

func (c *ChocolateDecorator) GetCost() float64 {
	return c.beverage.GetCost() + 15.0 // + стоимость шоколада
}
func (c *ChocolateDecorator) GetDescription() string {
	return c.beverage.GetDescription() + ", Chocolate"
}

// --- Задание 2: Дополнительные декораторы ---

// VanillaDecorator (Ваниль)
type VanillaDecorator struct {
	beverage IBeverage
}

func NewVanillaDecorator(b IBeverage) *VanillaDecorator {
	return &VanillaDecorator{beverage: b}
}

func (v *VanillaDecorator) GetCost() float64 {
	return v.beverage.GetCost() + 7.0 // + стоимость ванили
}
func (v *VanillaDecorator) GetDescription() string {
	return v.beverage.GetDescription() + ", Vanilla"
}

// CinnamonDecorator (Корица)
type CinnamonDecorator struct {
	beverage IBeverage
}

func NewCinnamonDecorator(b IBeverage) *CinnamonDecorator {
	return &CinnamonDecorator{beverage: b}
}

func (c *CinnamonDecorator) GetCost() float64 {
	return c.beverage.GetCost() + 6.0 // + стоимость корицы
}
func (c *CinnamonDecorator) GetDescription() string {
	return c.beverage.GetDescription() + ", Cinnamon"
}