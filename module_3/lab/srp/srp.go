package srp

import "fmt"

type Item struct {
	Name  string
	Price float64
}

type Invoice struct {
	ID      int
	Items   []Item
	TaxRate float64
}

type InvoiceCalculator struct{}

func (ic InvoiceCalculator) CalculateTotal(invoice Invoice) float64 {
	subTotal := 0.0
	for _, item := range invoice.Items {
		subTotal += item.Price
	}
	return subTotal + (subTotal * invoice.TaxRate)
}

type InvoiceRepository struct{}

func (repo InvoiceRepository) SaveToDatabase(invoice Invoice) {
	fmt.Printf("Invoice #%d saved to database with %d items\n", invoice.ID, len(invoice.Items))
}
