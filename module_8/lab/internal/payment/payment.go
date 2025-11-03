package payment

import "fmt"

// 1. Целевой интерфейс IPaymentProcessor
type IPaymentProcessor interface {
	ProcessPayment(amount float64)
	RefundPayment(amount float64)
}

// 2. Внутренняя платежная система
type InternalPaymentProcessor struct{}

func (p *InternalPaymentProcessor) ProcessPayment(amount float64) {
	fmt.Printf("Processing payment of %.2f via internal system.\n", amount)
}
func (p *InternalPaymentProcessor) RefundPayment(amount float64) {
	fmt.Printf("Refunding payment of %.2f via internal system.\n", amount)
}

// --- 3. Сторонние (Адаптируемые) системы ---

// ExternalPaymentSystemA
type ExternalPaymentSystemA struct{}

// У этой системы свои, несовместимые методы
func (s *ExternalPaymentSystemA) MakePayment(amount float64) {
	fmt.Printf("Making payment of %.2f via External Payment System A.\n", amount)
}
func (s *ExternalPaymentSystemA) MakeRefund(amount float64) {
	fmt.Printf("Making refund of %.2f via External Payment System A.\n", amount)
}

// ExternalPaymentSystemB
type ExternalPaymentSystemB struct{}

// У этой системы ДРУГИЕ несовместимые методы
func (s *ExternalPaymentSystemB) SendPayment(amount float64) {
	fmt.Printf("Sending payment of %.2f via External Payment System B.\n", amount)
}
func (s *ExternalPaymentSystemB) ProcessRefund(amount float64) {
	fmt.Printf("Processing refund of %.2f via External Payment System B.\n", amount)
}

// --- 4. Адаптеры ---

// PaymentAdapterA (Адаптер для SystemA)
type PaymentAdapterA struct {
	// Адаптер "оборачивает" несовместимый сервис
	externalSystemA *ExternalPaymentSystemA
}

func NewPaymentAdapterA(externalSystemA *ExternalPaymentSystemA) *PaymentAdapterA {
	return &PaymentAdapterA{externalSystemA: externalSystemA}
}

// Адаптер реализует целевой интерфейс
func (a *PaymentAdapterA) ProcessPayment(amount float64) {
	// ... и "переводит" вызов на нужный метод
	a.externalSystemA.MakePayment(amount)
}
func (a *PaymentAdapterA) RefundPayment(amount float64) {
	a.externalSystemA.MakeRefund(amount)
}

// PaymentAdapterB (Адаптер для SystemB)
type PaymentAdapterB struct {
	externalSystemB *ExternalPaymentSystemB
}

func NewPaymentAdapterB(externalSystemB *ExternalPaymentSystemB) *PaymentAdapterB {
	return &PaymentAdapterB{externalSystemB: externalSystemB}
}

// Адаптер B делает то же самое, но для методов SystemB
func (a *PaymentAdapterB) ProcessPayment(amount float64) {
	a.externalSystemB.SendPayment(amount)
}
func (a *PaymentAdapterB) RefundPayment(amount float64) {
	a.externalSystemB.ProcessRefund(amount)
}

// --- Дополнительное Задание 3: Логика выбора платежной системы ---

// Это "фабрика", которая решает, какой процессор выдать клиенту
// В реальном приложении сервисы (SystemA, SystemB) были бы синглтонами.
func GetPaymentProcessor(region string, currency string) IPaymentProcessor {
	fmt.Printf("\n[Фабрика: Получен запрос для Region: %s, Currency: %s]\n", region, currency)
	
	if region == "EU" && currency == "EUR" {
		fmt.Println("[Фабрика: Выбран ExternalSystemB для Европы/EUR]")
		// Создаем и возвращаем адаптер для SystemB
		return NewPaymentAdapterB(new(ExternalPaymentSystemB))
	}
	
	if region == "US" {
		fmt.Println("[Фабрика: Выбран ExternalSystemA для США]")
		// Создаем и возвращаем адаптер для SystemA
		return NewPaymentAdapterA(new(ExternalPaymentSystemA))
	}

	fmt.Println("[Фабрика: Выбрана внутренняя система по умолчанию]")
	// Во всех остальных случаях используем нашу внутреннюю систему
	return new(InternalPaymentProcessor)
} 