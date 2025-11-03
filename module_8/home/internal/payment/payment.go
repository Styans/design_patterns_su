package payment

import "fmt"

// 1. Целевой интерфейс
type IPaymentProcessor interface {
	ProcessPayment(amount float64)
}

// 2. Существующая реализация
type PayPalPaymentProcessor struct{}

func (p *PayPalPaymentProcessor) ProcessPayment(amount float64) {
	fmt.Printf("Processing PayPal payment of $%.2f\n", amount)
}

// --- 3. Сторонние (Адаптируемые) классы ---

// Stripe
type StripePaymentService struct{}

func (s *StripePaymentService) MakeTransaction(totalAmount float64) {
	fmt.Printf("Stripe transaction successful for $%.2f\n", totalAmount)
}

// CryptoGateway (Доп. задание)
type CryptoGateway struct{}

func (c *CryptoGateway) SubmitCryptoPayment(value float64, currency string) {
	fmt.Printf("Processing %f %s via CryptoGateway\n", value, currency)
}

// --- 4. Адаптеры ---

// Адаптер для Stripe
type StripePaymentAdapter struct {
	stripeService *StripePaymentService // неэкспортируемое поле
}

func NewStripePaymentAdapter(service *StripePaymentService) *StripePaymentAdapter {
	return &StripePaymentAdapter{stripeService: service}
}

// Адаптер реализует целевой интерфейс
func (a *StripePaymentAdapter) ProcessPayment(amount float64) {
	fmt.Println("Adapter converting ProcessPayment call to MakeTransaction call...")
	a.stripeService.MakeTransaction(amount)
}

// Адаптер для CryptoGateway (Доп. задание)
type CryptoGatewayAdapter struct {
	cryptoGateway *CryptoGateway // неэкспортируемое поле
}

func NewCryptoGatewayAdapter(gateway *CryptoGateway) *CryptoGatewayAdapter {
	return &CryptoGatewayAdapter{cryptoGateway: gateway}
}

// Адаптер реализует целевой интерфейс
func (a *CryptoGatewayAdapter) ProcessPayment(amount float64) {
	fmt.Println("Adapter converting ProcessPayment (USD) to SubmitCryptoPayment (BTC)...")
	btcAmount := amount / 40000.0 // Примерная конвертация
	a.cryptoGateway.SubmitCryptoPayment(btcAmount, "BTC")
}