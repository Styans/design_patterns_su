package main

import (
	"home/dip"
	"home/isp"
	"home/ocp"
	"home/srp"
	"fmt"
)

func main() {
	fmt.Println("=== Принципы проектирования — Стас ===")

	order := srp.Order{"Laptop", 2, 500.0}
	calculator := srp.OrderCalculator{}
	payment := srp.PaymentProcessor{}
	email := srp.EmailService{}

	total := calculator.CalculateTotalPrice(order)
	fmt.Println("SRP: Total price =", total)
	payment.ProcessPayment("Kaspi QR")
	email.SendConfirmationEmail("stas@example.com")

	permanent := ocp.PermanentEmployee{ocp.Employee{"Alex", 1000}}
	contract := ocp.ContractEmployee{ocp.Employee{"Tim", 900}}
	intern := ocp.Intern{ocp.Employee{"Nina", 700}}

	fmt.Println("OCP: Permanent salary =", permanent.CalculateSalary())
	fmt.Println("OCP: Contract salary =", contract.CalculateSalary())
	fmt.Println("OCP: Intern salary =", intern.CalculateSalary())

	allInOne := isp.AllInOnePrinter{}
	allInOne.Print("Document.pdf")
	allInOne.Scan("Photo.jpg")

	basic := isp.BasicPrinter{}
	basic.Print("Simple doc")

	emailSender := dip.EmailSender{}
	smsSender := dip.SmsSender{}
	service := dip.NewNotificationService(emailSender, smsSender)
	service.SendNotification("Design Principles test message")
}
