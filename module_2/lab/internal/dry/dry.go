package dry

import "fmt"

type OrderService struct{}

func (o OrderService) processOrder(action, productName string, quantity int, price float64) {
	total := float64(quantity) * price
	fmt.Printf("Order for %s %s. Total: %.2f\n", productName, action, total)
}

func (o OrderService) CreateOrder(productName string, quantity int, price float64) {
	o.processOrder("created", productName, quantity, price)
}

func (o OrderService) UpdateOrder(productName string, quantity int, price float64) {
	o.processOrder("updated", productName, quantity, price)
}

type Vehicle struct {
	Type string
}

func (v Vehicle) Start() {
	fmt.Printf("%s is starting\n", v.Type)
}

func (v Vehicle) Stop() {
	fmt.Printf("%s is stopping\n", v.Type)
}

type Car struct {
	Vehicle
}

type Truck struct {
	Vehicle
}
