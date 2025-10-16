package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"travelstocks/internal/stocks"
	"travelstocks/internal/travel"
)

func main() {
	demo := flag.String("demo", "travel", "travel|stocks")
	flag.Parse()
	switch *demo {
	case "travel":
		runTravel()
	case "stocks":
		runStocks()
	default:
		fmt.Println("travel|stocks")
	}
}

func runTravel() {
	req := travel.TripRequest{
		Transport:  "plane",
		Class:      travel.Business,
		Passengers: 3,
		Legs: []travel.Leg{
			{From: "ALA", To: "IST", DistanceKm: 3800, Region: "ASIA"},
			{From: "IST", To: "FRA", DistanceKm: 1900, Region: "EU"},
		},
		Discount: travel.PromoPercent,
		PromoPct: 12,
		Extras:   travel.Extras{CheckedBagsPerPax: 1, PriorityBoarding: true, Meal: true, WiFi: true, Insurance: true},
	}
	plane := travel.PlaneStrategy{}
	train := travel.TrainStrategy{}
	bus := travel.BusStrategy{}
	ctx := travel.NewTravelBookingContext(plane)
	total, breakdown, err := ctx.Quote(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(breakdown)
	fmt.Printf("PLANE TOTAL: %.2f\n\n", total)
	req.Transport = "train"
	ctx.SetStrategy(train)
	total, breakdown, _ = ctx.Quote(req)
	fmt.Println(breakdown)
	fmt.Printf("TRAIN TOTAL: %.2f\n\n", total)
	req.Transport = "bus"
	req.Class = travel.Economy
	req.Legs = []travel.Leg{{From: "ALA", To: "TAL", DistanceKm: 250, Region: "KZ"}}
	req.Discount = travel.Child
	req.Extras = travel.Extras{CheckedBagsPerPax: 0, Meal: true, WiFi: true}
	ctx.SetStrategy(bus)
	total, breakdown, _ = ctx.Quote(req)
	fmt.Println(breakdown)
	fmt.Printf("BUS TOTAL: %.2f\n", total)
}

func runStocks() {
	logger := log.New(log.Writer(), "[EXCH] ", log.LstdFlags|log.Lmicroseconds)
	ex := stocks.NewStockExchange(true, logger)
	tr1 := stocks.NewTrader("Alice", []string{"AAPL", "TSLA"})
	tr1.MinPrice["AAPL"] = 150
	tr2 := stocks.NewTrader("Bob", []string{"AAPL", "NVDA", "TSLA"})
	tr2.MaxPrice["TSLA"] = 350
	ro := stocks.NewRobot("R-1", []string{"AAPL", "TSLA"}, 100000, 10)
	ro.BuyBelow["AAPL"] = 155
	ro.SellAbove["AAPL"] = 190
	ro.BuyBelow["TSLA"] = 240
	ro.SellAbove["TSLA"] = 310
	ex.Subscribe(tr1, tr1.SubscribedTo()...)
	ex.Subscribe(tr2, tr2.SubscribedTo()...)
	ex.Subscribe(ro, ro.SubscribedTo()...)
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
		ex.SetPrice(s, stocks.Round2(p))
		time.Sleep(100 * time.Millisecond)
	}
	ex.Unsubscribe(tr2, "TSLA")
	for i := 0; i < 5; i++ {
		for _, s := range syms {
			pr[s] += (rand.Float64() - 0.5) * 6
			ex.SetPrice(s, stocks.Round2(math.Max(1, pr[s])))
			time.Sleep(80 * time.Millisecond)
		}
	}
	fmt.Println()
	fmt.Println(ex.Report())
	ex.Shutdown()
}
