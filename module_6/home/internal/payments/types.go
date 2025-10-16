package payments

import "time"

type PaymentRequest struct {
	Amount    float64
	Currency  string
	Customer  string
	Meta      map[string]string
	Timestamp time.Time
}

func (r PaymentRequest) WithAmount(a float64) PaymentRequest {
	r.Amount = a
	return r
}

type PaymentStrategy interface {
	Name() string
	Pay(PaymentRequest) string
}
