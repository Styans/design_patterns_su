package order

type Item struct {
	Name     string
	Quantity int
	Price    float64
}

type Order struct {
	ID           int
	Items        []Item
	PaymentName  string
	DeliveryName string
	CustomerEmail string
	Status       string
}

func (o *Order) AddItem(name string, qty int, price float64) {
	o.Items = append(o.Items, Item{Name: name, Quantity: qty, Price: price})
}

func (o Order) Subtotal() float64 {
	sum := 0.0
	for _, it := range o.Items {
		sum += float64(it.Quantity) * it.Price
	}
	return sum
}
