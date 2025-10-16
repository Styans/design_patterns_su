package strategies

import (
	"fmt"

	"payobs/internal/payments"
)

type card struct {
	num    string
	holder string
}

func NewCard(num, holder string) payments.PaymentStrategy {
	return &card{num: num, holder: holder}
}

func (c *card) Name() string { return "Card" }

func (c *card) Pay(req payments.PaymentRequest) string {
	m := "****"
	if len(c.num) >= 4 {
		m = "**** **** **** " + c.num[len(c.num)-4:]
	}
	fee := round2(req.Amount*0.015 + 50)
	total := round2(req.Amount + fee)
	return fmt.Sprintf("[%s]\namount=%.2f %s\ncard=%s holder=%s\nfee=1.5%%+50.00\ntotal=%.2f %s\n",
		c.Name(), req.Amount, req.Currency, m, c.holder, total, req.Currency)
}
