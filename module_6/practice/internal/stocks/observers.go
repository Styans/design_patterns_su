package stocks

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type Trader struct {
	id       string
	symbols  map[string]struct{}
	MinPrice map[string]float64
	MaxPrice map[string]float64
}

func NewTrader(id string, subs []string) *Trader {
	m := make(map[string]struct{})
	for _, s := range subs {
		m[s] = struct{}{}
	}
	return &Trader{id: id, symbols: m, MinPrice: map[string]float64{}, MaxPrice: map[string]float64{}}
}

func (t *Trader) ID() string { return t.id }

func (t *Trader) OnQuote(symbol string, price float64, ts time.Time) {
	fmt.Printf("[%s] %s @ %.2f (%s)\n", t.id, symbol, price, ts.Format(time.Kitchen))
}

func (t *Trader) Filter(symbol string, price float64) bool {
	if _, ok := t.symbols[symbol]; !ok {
		return false
	}
	if mn, ok := t.MinPrice[symbol]; ok && price < mn {
		return false
	}
	if mx, ok := t.MaxPrice[symbol]; ok && price > mx {
		return false
	}
	return true
}

func (t *Trader) SubscribedTo() []string {
	out := make([]string, 0, len(t.symbols))
	for s := range t.symbols {
		out = append(out, s)
	}
	sort.Strings(out)
	return out
}

type Robot struct {
	id         string
	watch      map[string]struct{}
	BuyBelow   map[string]float64
	SellAbove  map[string]float64
	position   map[string]float64
	cash       float64
	tradeSize  float64
	lastAction map[string]string
	mutex      sync.Mutex
}

func NewRobot(id string, subs []string, cash, tradeSize float64) *Robot {
	w := make(map[string]struct{})
	for _, s := range subs {
		w[s] = struct{}{}
	}
	return &Robot{
		id:         id,
		watch:      w,
		BuyBelow:   map[string]float64{},
		SellAbove:  map[string]float64{},
		position:   map[string]float64{},
		cash:       cash,
		tradeSize:  tradeSize,
		lastAction: map[string]string{},
	}
}

func (r *Robot) ID() string { return r.id }

func (r *Robot) Filter(symbol string, price float64) bool {
	_, ok := r.watch[symbol]
	return ok
}

func (r *Robot) SubscribedTo() []string {
	out := make([]string, 0, len(r.watch))
	for s := range r.watch {
		out = append(out, s)
	}
	sort.Strings(out)
	return out
}

func (r *Robot) OnQuote(symbol string, price float64, ts time.Time) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if th, ok := r.BuyBelow[symbol]; ok && price < th && r.cash >= r.tradeSize*price {
		r.position[symbol] += r.tradeSize
		r.cash -= r.tradeSize * price
		r.lastAction[symbol] = fmt.Sprintf("BUY %.0f @ %.2f", r.tradeSize, price)
		fmt.Printf("[ROBOT %s] BUY %s: %.0f @ %.2f | cash=%.2f pos=%.0f\n", r.id, symbol, r.tradeSize, price, r.cash, r.position[symbol])
		return
	}
	if th, ok := r.SellAbove[symbol]; ok && price > th && r.position[symbol] >= r.tradeSize {
		r.position[symbol] -= r.tradeSize
		r.cash += r.tradeSize * price
		r.lastAction[symbol] = fmt.Sprintf("SELL %.0f @ %.2f", r.tradeSize, price)
		fmt.Printf("[ROBOT %s] SELL %s: %.0f @ %.2f | cash=%.2f pos=%.0f\n", r.id, symbol, r.tradeSize, price, r.cash, r.position[symbol])
		return
	}
}
