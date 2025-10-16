package strategies

import (
	"fmt"
	"math"

	"payobs/internal/payments"
)

type paypal struct {
	email string
}

func NewPayPal(email string) payments.PaymentStrategy {
	return &paypal{email: email}
}

func (p *paypal) Name() string { return "PayPal" }

func (p *paypal) Pay(req payments.PaymentRequest) string {
	fee := round2(math.Max(0.035*req.Amount, 60))
	total := round2(req.Amount + fee)
	return fmt.Sprintf("[%s]\namount=%.2f %s\naccount=%s\nfee=max(3.5%%,60)\ntotal=%.2f %s\n",
		p.Name(), req.Amount, req.Currency, p.email, total, req.Currency)
}
