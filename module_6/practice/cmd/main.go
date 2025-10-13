package main

import "fmt"

func main() {
	airplane := &AirplaneCost{}
	train := &TrainCost{}
	bus := &BusCost{}
	distance := 0.0
	passanger := 2
	serviceClass := "business"
	// bus.CalculateCost(distance, passanger, serviceClass)

	arrayServiceClass := []string{"economy", "business", "first"}
	for _, v := range arrayServiceClass {
		serviceClass = v
		fmt.Printf("Service Class: %s\n", serviceClass)
		distance = 1000.0

		fmt.Printf("Airplane Cost: %.2f\n",
			airplane.CalculateCost(distance, passanger, serviceClass),
		)

		fmt.Printf("Train Cost: %.2f\n",
			train.CalculateCost(distance, passanger, serviceClass),
		)

		fmt.Printf("Bus Cost: %.2f\n",
			bus.CalculateCost(distance, passanger, serviceClass),
		)

		distance = 2000.0
		fmt.Println()
	}
	// airplaneCost := CalculateTravelCost(airplane, distance, passanger, serviceClass)
	// trainCost := CalculateTravelCost(train, distance, passanger, serviceClass)
	// busCost := CalculateTravelCost(bus, distance, passanger, serviceClass)
	// fmt.Printf("Airplane Cost: %.2f\n", airplaneCost)
	// fmt.Printf("Train Cost: %.2f\n", trainCost)
	// fmt.Printf("Bus Cost: %.2f\n", busCost)

	Trader1 := &Trader{}
	Trader2 := &Trader{}
	stockExchange := &StockExchange{}
	stockExchange.RegisterObserver("AAPL", Trader1)
	stockExchange.RegisterObserver("GOOGL", Trader1)
	stockExchange.RegisterObserver("AAPL", Trader2)
	stockExchange.prices = map[string]float64{
		"AAPL":  150.00,
		"GOOGL": 2800.00,
	}
	stockExchange.NotifyObservers("AAPL")
	stockExchange.NotifyObservers("GOOGL")
	stockExchange.prices["AAPL"] = 155.00
	stockExchange.NotifyObservers("AAPL")
	stockExchange.RemoveObserver("AAPL", Trader2)
	stockExchange.prices["AAPL"] = 160.00
	stockExchange.NotifyObservers("AAPL")
	fmt.Println()

}

type ICostCalculateStrategy interface {
	CalculateCost(amount float64, passanger int, serviceClass string) float64
}

type AirplaneCost struct{}

func (a *AirplaneCost) CalculateCost(distance float64, passanger int, serviceClass string) float64 {
	basecost := 0.5 * distance
	if serviceClass == "business" {
		basecost *= 1.5
	}
	if serviceClass == "first" {
		basecost *= 2
	}

	return basecost*float64(passanger) + 50
}

type TrainCost struct{}

func (t *TrainCost) CalculateCost(distance float64, passanger int, serviceClass string) float64 {
	basecost := 0.2 * distance
	if serviceClass == "business" {
		basecost *= 1.3
	}
	if serviceClass == "first" {
		basecost *= 1.7
	}
	return basecost*float64(passanger) + 20
}

type BusCost struct{}

func (b *BusCost) CalculateCost(distance float64, passanger int, serviceClass string) float64 {
	basecost := 0.1 * distance
	if serviceClass == "business" {
		basecost *= 1.2
	}
	if serviceClass == "first" {
		basecost *= 1.5
	}
	return basecost*float64(passanger) + 10
}

type IObserver interface {
	Update(stockSymbol string, price float64)
}

type Isubject interface {
	RegisterObserver(stockSymbol string, observer IObserver)
	RemoveObserver(stockSymbol string, observer IObserver)
	NotifyObservers(stockSymbol string)
}

type StockExchange struct {
	observers map[string]IObserver
	prices    map[string]float64
}

func (se *StockExchange) RegisterObserver(stockSymbol string, observer IObserver) {

	if se.observers[stockSymbol] != nil {

		se.observers[stockSymbol] = observer
	}

}

func (se *StockExchange) RemoveObserver(stockSymbol string, observer IObserver) {
	if se.observers == nil || se.observers[stockSymbol] == nil {
		return
	}
	delete(se.observers[stockSymbol], observer)
	if len(se.observers[stockSymbol]) == 0 {
		delete(se.observers, stockSymbol)
	}
}
func (se *StockExchange) NotifyObservers(stockSymbol string) {
	if se.observers == nil || se.observers[stockSymbol] == nil {
		return
	}
	price, exists := se.prices[stockSymbol]
	if !exists {
		return
	}
	for observer := range se.observers[stockSymbol] {
		fmt.Println(stockSymbol)
		observer.Update(stockSymbol, price)
	}
}

type Trader struct{}

func (t *Trader) Update(stockSymbol string, price float64) {
	fmt.Printf("Trader notified: %s price changed to %.2f\n", stockSymbol, price)
}
