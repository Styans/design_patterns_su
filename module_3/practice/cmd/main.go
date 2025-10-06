package main

import (
	"fmt"
	"practice/internal/discount"
	"practice/internal/delivery"
	"practice/internal/notification"
	"practice/internal/order"
	"practice/internal/payment"
)

type OrderProcessor struct {
	payment    payment.Payment
	delivery   delivery.Delivery
	notifiers  []notification.Notifier
	calculator discount.DiscountCalculator
}

func NewOrderProcessor(p payment.Payment, d delivery.Delivery, calc discount.DiscountCalculator, notifiers ...notification.Notifier) OrderProcessor {
	return OrderProcessor{payment: p, delivery: d, calculator: calc, notifiers: notifiers}
}

func (op OrderProcessor) Process(o *order.Order) error {
	sub := o.Subtotal()
	total := op.calculator.Calculate(sub)
	err := op.payment.ProcessPayment(total)
	if err != nil {
		o.Status = "PaymentFailed"
		op.notifyAll("Payment failed for order")
		return err
	}
	o.Status = "Paid"
	err = op.delivery.DeliverOrder(*o)
	if err != nil {
		o.Status = "DeliveryFailed"
		op.notifyAll("Delivery failed for order")
		return err
	}
	o.Status = "Completed"
	op.notifyAll("Order completed")
	return nil
}

func (op OrderProcessor) notifyAll(msg string) {
	for _, n := range op.notifiers {
		_ = n.SendNotification(msg)
	}
}

func main() {
	o := order.Order{ID: 1, CustomerEmail: "stas@example.com"}
	o.AddItem("Ticket", 2, 50)
	o.AddItem("T-shirt", 1, 20)

	disc := discount.NewCalculator(discount.PercentageDiscount{Percent: 10})

	p := payment.CreditCardPayment{CardNumber: "1234-****-****-5678"}

	d := delivery.CourierDelivery{Address: "Almaty, Dostyk 1"}

	email := notification.EmailNotification{To: o.CustomerEmail}
	sms := notification.SmsNotification{Phone: "+77000000000"}

	processor := NewOrderProcessor(p, d, disc, email, sms)

	fmt.Println("Order subtotal:", o.Subtotal())
	total := disc.Calculate(o.Subtotal())
	fmt.Println("Order total after discounts:", total)

	err := processor.Process(&o)
	if err != nil {
		fmt.Println("Processing error:", err)
	} else {
		fmt.Println("Order status:", o.Status)
	}

	fmt.Println("Payment method:", p.Name())
	fmt.Println("Delivery method:", d.Name())
	for _, n := range []notification.Notifier{email, sms} {
		fmt.Println("Notifier:", n.Name())
	}
}
