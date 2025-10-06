package main

import (
	"fmt"
	"lab/dip"
	"lab/isp"
	"lab/ocp"
	"lab/srp"
)

func main() {
	fmt.Println("=== Лабораторная работа: Принципы проектирования (Стас) ===\n")

	fmt.Println("SRP Example:")
	invoice := srp.Invoice{
		ID: 1,
		Items: []srp.Item{
			{"Laptop", 500},
			{"Mouse", 50},
		},
		TaxRate: 0.12,
	}
	calculator := srp.InvoiceCalculator{}
	repo := srp.InvoiceRepository{}

	total := calculator.CalculateTotal(invoice)
	fmt.Printf("Total amount with tax: %.2f\n", total)
	repo.SaveToDatabase(invoice)
	fmt.Println()

	fmt.Println("OCP Example:")
	gold := ocp.GoldCustomer{}
	calculatorOCP := ocp.NewDiscountCalculator(gold)
	fmt.Printf("Gold discount total: %.2f\n", calculatorOCP.Calculate(1000))
	fmt.Println()

	fmt.Println("ISP Example:")
	human := isp.HumanWorker{}
	robot := isp.RobotWorker{}
	human.Work()
	human.Eat()
	human.Sleep()
	robot.Work()
	fmt.Println()

	fmt.Println("DIP Example:")
	emailNotifier := dip.EmailService{}
	smsNotifier := dip.SmsService{}

	emailNotification := dip.NewNotification(emailNotifier)
	smsNotification := dip.NewNotification(smsNotifier)

	emailNotification.Send("Привет, Стас! Это уведомление по email.")
	smsNotification.Send("Привет, Стас! Это уведомление по SMS.")
}
