package srp

import "fmt"

type Order struct {
	ProductName string
	Quantity    int
	Price       float64
}

type OrderCalculator struct{}

func (c OrderCalculator) CalculateTotalPrice(order Order) float64 {
	return order.Price * float64(order.Quantity) * 0.9
}

type PaymentProcessor struct{}

func (p PaymentProcessor) ProcessPayment(paymentDetails string) {
	fmt.Println("Payment processed using:", paymentDetails)
}

type EmailService struct{}

func (e EmailService) SendConfirmationEmail(email string) {
	fmt.Println("Confirmation email sent to:", email)
}
