package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"payobs/internal/observer"
	"payobs/internal/payments"
	"payobs/internal/payments/strategies"
)

func main() {
	mode := flag.String("mode", "payments", "payments|observer")
	flag.Parse()
	if *mode == "payments" {
		runPayments()
		return
	}
	runObserver()
}

func runPayments() {
	ctx := payments.NewContext()
	req := payments.PaymentRequest{
		Amount:    9999.90,
		Currency:  "KZT",
		Customer:  "CUST-001",
		Meta:      map[string]string{"orderId": "ORD-1001"},
		Timestamp: time.Now(),
	}
	ctx.SetStrategy(strategies.NewCard("5555444433331111", "ZHANERKE AKAN"))
	fmt.Println(ctx.Pay(req))
	ctx.SetStrategy(strategies.NewPayPal("user@example.com"))
	fmt.Println(ctx.Pay(req.WithAmount(12500)))
	ctx.SetStrategy(strategies.NewCrypto("0xAE12B3CD45EF99887766", "USDT", 0.01))
	fmt.Println(ctx.Pay(req.WithAmount(7300)))

	in := bufio.NewReader(os.Stdin)
	fmt.Print("method [card|paypal|crypto]: ")
	choice, _ := in.ReadString('\n')
	choice = strings.TrimSpace(strings.ToLower(choice))
	var strat payments.PaymentStrategy
	switch choice {
	case "card":
		strat = strategies.NewCard("4111111111111111", "JANE DOE")
	case "paypal":
		strat = strategies.NewPayPal("buyer@example.com")
	case "crypto":
		strat = strategies.NewCrypto("0xBEEFCAFE1234", "BTC", 0.015)
	default:
		strat = strategies.NewCard("4111111111111111", "JANE DOE")
	}
	ctx.SetStrategy(strat)
	fmt.Print("amount: ")
	a, _ := in.ReadString('\n')
	a = strings.TrimSpace(a)
	val, err := strconv.ParseFloat(strings.ReplaceAll(a, ",", "."), 64)
	if err != nil || val <= 0 {
		fmt.Println("invalid amount")
		return
	}
	fmt.Println(ctx.Pay(req.WithAmount(val)))
}

func runObserver() {
	ex := observer.NewStockExchange(true, nil)
	o1 := observer.NewConsoleObserver("Alice", []string{"AAPL", "TSLA"})
	o2 := observer.NewThresholdObserver("Bot-1", []string{"AAPL", "NVDA"}, map[string]float64{"AAPL": 155, "TSLA": 240}, map[string]float64{"AAPL": 190, "TSLA": 310})
	o3 := observer.NewStatsObserver()
	ex.Subscribe(o1, o1.SubscribedTo()...)
	ex.Subscribe(o2, o2.SubscribedTo()...)
	ex.Subscribe(o3, "AAPL", "TSLA", "NVDA")
	syms := []string{"AAPL", "TSLA", "NVDA"}
	pr := map[string]float64{"AAPL": 160, "TSLA": 280, "NVDA": 460}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		s := syms[rand.Intn(len(syms))]
		if i == 10 {
			pr["AAPL"] += 25
		} else {
			pr[s] += (rand.Float64() - 0.5) * 10
		}
		p := math.Max(1, pr[s])
		ex.SetPrice(s, round2(p))
		time.Sleep(100 * time.Millisecond)
	}
	ex.Unsubscribe(o1, "TSLA")
	for i := 0; i < 5; i++ {
		for _, s := range syms {
			pr[s] += (rand.Float64() - 0.5) * 6
			ex.SetPrice(s, round2(math.Max(1, pr[s])))
			time.Sleep(80 * time.Millisecond)
		}
	}
	fmt.Println()
	fmt.Println(ex.Report())
	ex.Shutdown()
}

func round2(x float64) float64 { return math.Round(x*100) / 100 }
