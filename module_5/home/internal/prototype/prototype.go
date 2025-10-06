package prototype

type Cloneable interface {
	Clone() Cloneable
}

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func (p *Product) Clone() Cloneable {
	copy := *p
	return &copy
}

type Discount struct {
	Description string
	Percent     float64
}

func (d *Discount) Clone() Cloneable {
	copy := *d
	return &copy
}

type Order struct {
	Products []Product
	Delivery float64
	Discount Discount
	Payment  string
}

func (o *Order) Clone() Cloneable {
	newProducts := make([]Product, len(o.Products))
	for i, p := range o.Products {
		newProducts[i] = *p.Clone().(*Product)
	}
	newDiscount := *o.Discount.Clone().(*Discount)
	copy := *o
	copy.Products = newProducts
	copy.Discount = newDiscount
	return &copy
}
