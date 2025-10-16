package observer

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)

type quote struct {
	symbol string
	price  float64
	ts     time.Time
}

type StockExchange struct {
	mu           sync.RWMutex
	prices       map[string]float64
	subs         map[string]map[string]Observer
	async        bool
	wg           sync.WaitGroup
	queue        chan quote
	notifyCounts map[string]int
}

func NewStockExchange(async bool, _ interface{}) *StockExchange {
	se := &StockExchange{
		prices:       map[string]float64{},
		subs:         map[string]map[string]Observer{},
		async:        async,
		queue:        make(chan quote, 256),
		notifyCounts: map[string]int{},
	}
	if async {
		go se.dispatch()
	}
	return se
}

func (se *StockExchange) dispatch() {
	for q := range se.queue {
		se.mu.RLock()
		os := se.subs[q.symbol]
		se.mu.RUnlock()
		for _, o := range os {
			if !o.Filter(q.symbol, q.price) {
				continue
			}
			se.wg.Add(1)
			go func(obs Observer) {
				defer se.wg.Done()
				obs.OnQuote(q.symbol, q.price, q.ts)
				se.mu.Lock()
				se.notifyCounts[obs.ID()]++
				se.mu.Unlock()
			}(o)
		}
	}
}

func (se *StockExchange) Subscribe(obs Observer, symbols ...string) {
	se.mu.Lock()
	defer se.mu.Unlock()
	for _, s := range symbols {
		if se.subs[s] == nil {
			se.subs[s] = map[string]Observer{}
		}
		se.subs[s][obs.ID()] = obs
	}
}

func (se *StockExchange) Unsubscribe(obs Observer, symbols ...string) {
	se.mu.Lock()
	defer se.mu.Unlock()
	if len(symbols) == 0 {
		for s := range se.subs {
			delete(se.subs[s], obs.ID())
		}
		return
	}
	for _, s := range symbols {
		delete(se.subs[s], obs.ID())
	}
}

func (se *StockExchange) SetPrice(symbol string, price float64) {
	ts := time.Now()
	se.mu.Lock()
	se.prices[symbol] = price
	se.mu.Unlock()
	if se.async {
		se.queue <- quote{symbol: symbol, price: price, ts: ts}
		return
	}
	se.mu.RLock()
	os := se.subs[symbol]
	se.mu.RUnlock()
	for _, o := range os {
		if !o.Filter(symbol, price) {
			continue
		}
		o.OnQuote(symbol, price, ts)
		se.mu.Lock()
		se.notifyCounts[o.ID()]++
		se.mu.Unlock()
	}
}

func (se *StockExchange) ListSubscriptions() map[string][]string {
	se.mu.RLock()
	defer se.mu.RUnlock()
	out := map[string][]string{}
	for sym, m := range se.subs {
		for id := range m {
			out[sym] = append(out[sym], id)
		}
		sort.Strings(out[sym])
	}
	return out
}

func (se *StockExchange) Report() string {
	se.mu.RLock()
	defer se.mu.RUnlock()
	var b strings.Builder
	b.WriteString("=== Subscriptions ===\n")
	keys := make([]string, 0, len(se.subs))
	for k := range se.subs {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, s := range keys {
		var ids []string
		for id := range se.subs[s] {
			ids = append(ids, id)
		}
		sort.Strings(ids)
		b.WriteString(fmt.Sprintf("%s -> %v\n", s, ids))
	}
	b.WriteString("\n=== Notify counts ===\n")
	ids := make([]string, 0, len(se.notifyCounts))
	for id := range se.notifyCounts {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	for _, id := range ids {
		b.WriteString(fmt.Sprintf("%s: %d\n", id, se.notifyCounts[id]))
	}
	return b.String()
}

func (se *StockExchange) Shutdown() {
	if se.async {
		close(se.queue)
		se.wg.Wait()
	}
}
