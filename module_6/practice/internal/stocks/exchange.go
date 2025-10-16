package stocks

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"sync"
	"time"
)

type IObserver interface {
	ID() string
	OnQuote(symbol string, price float64, ts time.Time)
	Filter(symbol string, price float64) bool
	SubscribedTo() []string
}

type ISubject interface {
	Subscribe(obs IObserver, symbols ...string)
	Unsubscribe(obs IObserver, symbols ...string)
	SetPrice(symbol string, price float64)
	ListSubscriptions() map[string][]string
}

type StockExchange struct {
	mu           sync.RWMutex
	prices       map[string]float64
	subs         map[string]map[string]IObserver
	notifyAsync  bool
	notifyWG     sync.WaitGroup
	notifyQueue  chan quote
	logger       *log.Logger
	notifyCounts map[string]int
}

type quote struct {
	symbol string
	price  float64
	ts     time.Time
}

func NewStockExchange(async bool, logger *log.Logger) *StockExchange {
	se := &StockExchange{
		prices:       make(map[string]float64),
		subs:         make(map[string]map[string]IObserver),
		notifyAsync:  async,
		notifyQueue:  make(chan quote, 256),
		logger:       logger,
		notifyCounts: make(map[string]int),
	}
	if async {
		go se.dispatcher()
	}
	return se
}

func (se *StockExchange) dispatcher() {
	for q := range se.notifyQueue {
		se.mu.RLock()
		observers := se.subs[q.symbol]
		se.mu.RUnlock()
		for _, o := range observers {
			if !o.Filter(q.symbol, q.price) {
				continue
			}
			se.notifyWG.Add(1)
			go func(obs IObserver) {
				defer se.notifyWG.Done()
				obs.OnQuote(q.symbol, q.price, q.ts)
				se.mu.Lock()
				se.notifyCounts[obs.ID()]++
				se.mu.Unlock()
			}(o)
		}
	}
}

func (se *StockExchange) Subscribe(obs IObserver, symbols ...string) {
	se.mu.Lock()
	defer se.mu.Unlock()
	for _, s := range symbols {
		if se.subs[s] == nil {
			se.subs[s] = make(map[string]IObserver)
		}
		se.subs[s][obs.ID()] = obs
		if se.logger != nil {
			se.logger.Printf("Subscribe: %s -> %s\n", obs.ID(), s)
		}
	}
}

func (se *StockExchange) Unsubscribe(obs IObserver, symbols ...string) {
	se.mu.Lock()
	defer se.mu.Unlock()
	if len(symbols) == 0 {
		for s := range se.subs {
			delete(se.subs[s], obs.ID())
			if se.logger != nil {
				se.logger.Printf("Unsubscribe: %s -/-> %s\n", obs.ID(), s)
			}
		}
		return
	}
	for _, s := range symbols {
		delete(se.subs[s], obs.ID())
		if se.logger != nil {
			se.logger.Printf("Unsubscribe: %s -/-> %s\n", obs.ID(), s)
		}
	}
}

func (se *StockExchange) SetPrice(symbol string, price float64) {
	ts := time.Now()
	se.mu.Lock()
	se.prices[symbol] = price
	se.mu.Unlock()
	if se.logger != nil {
		se.logger.Printf("Price update: %s = %.2f\n", symbol, price)
	}
	if se.notifyAsync {
		se.notifyQueue <- quote{symbol, price, ts}
	} else {
		se.mu.RLock()
		observers := se.subs[symbol]
		se.mu.RUnlock()
		for _, o := range observers {
			if !o.Filter(symbol, price) {
				continue
			}
			o.OnQuote(symbol, price, ts)
			se.mu.Lock()
			se.notifyCounts[o.ID()]++
			se.mu.Unlock()
		}
	}
}

func (se *StockExchange) ListSubscriptions() map[string][]string {
	se.mu.RLock()
	defer se.mu.RUnlock()
	out := make(map[string][]string)
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
	var sb strings.Builder
	sb.WriteString("=== Subscriptions ===\n")
	keys := make([]string, 0, len(se.subs))
	for k := range se.subs {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, sym := range keys {
		var ids []string
		for id := range se.subs[sym] {
			ids = append(ids, id)
		}
		sort.Strings(ids)
		sb.WriteString(fmt.Sprintf("%s -> %v\n", sym, ids))
	}
	sb.WriteString("\n=== Notify counts ===\n")
	ids := make([]string, 0, len(se.notifyCounts))
	for id := range se.notifyCounts {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	for _, id := range ids {
		sb.WriteString(fmt.Sprintf("%s: %d\n", id, se.notifyCounts[id]))
	}
	return sb.String()
}

func (se *StockExchange) Shutdown() {
	if se.notifyAsync {
		close(se.notifyQueue)
		se.notifyWG.Wait()
	}
}
