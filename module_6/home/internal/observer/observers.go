package observer

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type ConsoleObserver struct {
	id      string
	symbols map[string]struct{}
}

func NewConsoleObserver(id string, subs []string) *ConsoleObserver {
	m := map[string]struct{}{}
	for _, s := range subs {
		m[s] = struct{}{}
	}
	return &ConsoleObserver{id: id, symbols: m}
}

func (c *ConsoleObserver) ID() string { return c.id }

func (c *ConsoleObserver) SubscribedTo() []string {
	out := make([]string, 0, len(c.symbols))
	for s := range c.symbols {
		out = append(out, s)
	}
	sort.Strings(out)
	return out
}

func (c *ConsoleObserver) Filter(symbol string, _ float64) bool {
	_, ok := c.symbols[symbol]
	return ok
}

func (c *ConsoleObserver) OnQuote(symbol string, price float64, ts time.Time) {
	fmt.Printf("[%s] %s @ %.2f (%s)\n", c.id, symbol, price, ts.Format(time.Kitchen))
}

type StatsObserver struct {
	id   string
	hist map[string][]float64
}

func NewStatsObserver() *StatsObserver {
	return &StatsObserver{id: "Stats", hist: map[string][]float64{}}
}

func (s *StatsObserver) ID() string { return s.id }

func (s *StatsObserver) SubscribedTo() []string { return []string{} }

func (s *StatsObserver) Filter(string, float64) bool { return true }

func (s *StatsObserver) OnQuote(symbol string, price float64, _ time.Time) {
	s.hist[symbol] = append(s.hist[symbol], price)
	avg := 0.0
	for _, v := range s.hist[symbol] {
		avg += v
	}
	avg /= float64(len(s.hist[symbol]))
	fmt.Printf("[Stats] %s: n=%d, avg=%.2f\n", symbol, len(s.hist[symbol]), avg)
}

type ThresholdObserver struct {
	id        string
	watch     map[string]struct{}
	buyBelow  map[string]float64
	sellAbove map[string]float64
	pos       map[string]float64
	cash      float64
	lot       float64
	mu        sync.Mutex
}

func NewThresholdObserver(id string, subs []string, buyBelow, sellAbove map[string]float64) *ThresholdObserver {
	w := map[string]struct{}{}
	for _, s := range subs {
		w[s] = struct{}{}
	}
	return &ThresholdObserver{
		id:        id,
		watch:     w,
		buyBelow:  buyBelow,
		sellAbove: sellAbove,
		pos:       map[string]float64{},
		cash:      100000,
		lot:       10,
	}
}

func (t *ThresholdObserver) ID() string { return t.id }

func (t *ThresholdObserver) SubscribedTo() []string {
	out := make([]string, 0, len(t.watch))
	for s := range t.watch {
		out = append(out, s)
	}
	sort.Strings(out)
	return out
}

func (t *ThresholdObserver) Filter(symbol string, _ float64) bool {
	_, ok := t.watch[symbol]
	return ok
}

func (t *ThresholdObserver) OnQuote(symbol string, price float64, _ time.Time) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if th, ok := t.buyBelow[symbol]; ok && price < th && t.cash >= t.lot*price {
		t.pos[symbol] += t.lot
		t.cash -= t.lot * price
		fmt.Printf("[ROBOT %s] BUY %s %.0f @ %.2f cash=%.2f pos=%.0f\n", t.id, symbol, t.lot, price, t.cash, t.pos[symbol])
		return
	}
	if th, ok := t.sellAbove[symbol]; ok && price > th && t.pos[symbol] >= t.lot {
		t.pos[symbol] -= t.lot
		t.cash += t.lot * price
		fmt.Printf("[ROBOT %s] SELL %s %.0f @ %.2f cash=%.2f pos=%.0f\n", t.id, symbol, t.lot, price, t.cash, t.pos[symbol])
		return
	}
}
