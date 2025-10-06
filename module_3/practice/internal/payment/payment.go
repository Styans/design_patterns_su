package payment

type Payment interface {
	ProcessPayment(amount float64) error
	Name() string
}

type CreditCardPayment struct {
	CardNumber string
}

func (c CreditCardPayment) ProcessPayment(amount float64) error {
	return nil
}

func (CreditCardPayment) Name() string {
	return "CreditCard"
}

type PayPalPayment struct {
	Account string
}

func (p PayPalPayment) ProcessPayment(amount float64) error {
	return nil
}

func (PayPalPayment) Name() string {
	return "PayPal"
}

type BankTransferPayment struct {
	Account string
}

func (b BankTransferPayment) ProcessPayment(amount float64) error {
	return nil
}

func (BankTransferPayment) Name() string {
	return "BankTransfer"
}
