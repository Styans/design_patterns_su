package main

import "fmt"

func main() {
	creditCard := &CreditCardPayment{cardNumber: "1234567812345678"}
	payPal := &PayPalPayment{email: "milka@gmail.com"}
	bitcoin := &BitcoinPayment{walletAddress: "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa"}
	Order(creditCard, 100.0)
	Order(payPal, 200.0)
	Order(bitcoin, 300.0)
	mainIObserver()
}

func mainIObserver() {
	eventManager := NewEventManager()
	emailNotifier := &EmailNotifier{email: "milka@gmail.com"}
	smsNotifier := &SMSNotifier{phoneNumber: "+1234567890"}
	eventManager.RegisterObserv(emailNotifier)
	eventManager.RegisterObserv(smsNotifier)
	eventManager.Notify("Order Placed")
	eventManager.RemoveObserv(smsNotifier)
	eventManager.Notify("Order Shipped")
	weatherDisplay := &CurrentConditionsDisplay{}
	eventManager.RegisterObserv(weatherDisplay)
	weatherDisplay.Update("72.5")

	eventManager.Notify("Order Shipped")

}

type IPaymentStrategy interface {
	Pay(amount float64) string
}

type CreditCardPayment struct {
	cardNumber string
}

type PayPalPayment struct {
	email string
}

type BitcoinPayment struct {
	walletAddress string
}

func (c *CreditCardPayment) Pay(amount float64) string {
	return "Paid " + formatAmount(amount) + " using Credit Card ending with " + c.cardNumber
}

func (p *PayPalPayment) Pay(amount float64) string {
	return "Paid " + formatAmount(amount) + " using PayPal account " + p.email
}

func (b *BitcoinPayment) Pay(amount float64) string {
	return "Paid " + formatAmount(amount) + " using Bitcoin wallet " + b.walletAddress
}

func formatAmount(amount float64) string {
	return "$" + fmt.Sprintf("%.2f", amount)
}

func Order(strategy IPaymentStrategy, amount float64) {
	result := strategy.Pay(amount)
	fmt.Println(result)
}

type IObserver interface {
	Update(event string)
}

type Subject interface {
	RegisterObserv(Iobserver IObserver)
	RemoveObserv(Iobserver IObserver)
	Notify(event string)
}

type EventManager struct {
	Iobservers map[IObserver]struct{}
}

func NewEventManager() *EventManager {
	return &EventManager{
		Iobservers: make(map[IObserver]struct{}),
	}
}

func (em *EventManager) RegisterObserv(Iobserver IObserver) {
	em.Iobservers[Iobserver] = struct{}{}
}

func (em *EventManager) RemoveObserv(Iobserver IObserver) {
	delete(em.Iobservers, Iobserver)
}

func (em *EventManager) Notify(event string) {
	for Iobserver := range em.Iobservers {
		Iobserver.Update(event)
	}
}

type EmailNotifier struct {
	email string
}

func (en *EmailNotifier) Update(event string) {
	fmt.Printf("Email to %s: New event - %s\n", en.email, event)
}

type SMSNotifier struct {
	phoneNumber string
}

func (sn *SMSNotifier) Update(event string) {
	fmt.Printf("SMS to %s: New event - %s\n", sn.phoneNumber, event)
}

type CurrentConditionsDisplay struct {
	temperature float64
	humidity    float64
}

func (ccd *CurrentConditionsDisplay) Update(temperature string) {
	ccd.Display()
}

func (ccd *CurrentConditionsDisplay) Display() {
	fmt.Printf("Current conditions: %.2f°F and %.2f%% humidity\n", ccd.temperature, ccd.humidity)
}
