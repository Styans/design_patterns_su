package observer

import "time"

type Observer interface {
	ID() string
	SubscribedTo() []string
	Filter(symbol string, price float64) bool
	OnQuote(symbol string, price float64, ts time.Time)
}

type Subject interface {
	Subscribe(obs Observer, symbols ...string)
	Unsubscribe(obs Observer, symbols ...string)
	SetPrice(symbol string, price float64)
	ListSubscriptions() map[string][]string
	Report() string
	Shutdown()
}
