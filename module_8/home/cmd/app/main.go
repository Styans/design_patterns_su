package main

import (
	"fmt"
	"strings"

	// Импортируем наши внутренние пакеты
	"cafe_system/internal/beverage"
	"cafe_system/internal/payment"
)

// ==================================================================
//
//	MAIN: Клиентский код для демонстрации
//
// ==================================================================
func main() {
	// --- Демонстрация "Декоратора" ---
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("   ЧАСТЬ 1: ПАТТЕРН \"ДЕКОРАТОР\"")
	fmt.Println(strings.Repeat("=", 50))

	// Тест 1: Простой эспрессо
	fmt.Println("Тест 1: Простой эспрессо")
	beverage1 := &beverage.Espresso{}
	printBeverageDetails(beverage1)

	// Тест 2: Чай с молоком и сахаром
	fmt.Println("Тест 2: Чай с молоком и сахаром")
	var beverage2 beverage.Beverage = &beverage.Tea{}
	beverage2 = beverage.NewMilk(beverage2)
	beverage2 = beverage.NewSugar(beverage2)
	printBeverageDetails(beverage2)

	// Тест 3: Мокка с взбитыми сливками, сиропом и корицей
	fmt.Println("Тест 3: Мокка с взбитыми сливками, сиропом и корицей")
	var beverage3 beverage.Beverage = &beverage.Mocha{}
	beverage3 = beverage.NewWhippedCream(beverage3)
	beverage3 = beverage.NewSyrup(beverage3)
	beverage3 = beverage.NewCinnamon(beverage3)
	printBeverageDetails(beverage3)

	// Тест 4: Латте с двойным молоком и сахаром
	fmt.Println("Тест 4: Латте с двойным молоком и сахаром")
	var beverage4 beverage.Beverage = &beverage.Latte{}
	beverage4 = beverage.NewMilk(beverage4)
	beverage4 = beverage.NewMilk(beverage4)
	beverage4 = beverage.NewSugar(beverage4)
	printBeverageDetails(beverage4)

	// --- Демонстрация "Адаптера" ---
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("   ЧАСТЬ 2: ПАТТЕРН \"АДАПТЕР\"")
	fmt.Println(strings.Repeat("=", 50))

	paymentAmount := 250.45

	// Создаем список процессоров, используя ЕДИНЫЙ интерфейс
	processors := []payment.IPaymentProcessor{}

	// 1. PayPal (стандартная реализация)
	payPal := &payment.PayPalPaymentProcessor{}
	processors = append(processors, payPal)

	// 2. Stripe (через Адаптер)
	stripeService := &payment.StripePaymentService{}
	stripeAdapter := payment.NewStripePaymentAdapter(stripeService)
	processors = append(processors, stripeAdapter)

	// 3. Crypto (через Адаптер)
	cryptoGateway := &payment.CryptoGateway{}
	cryptoAdapter := payment.NewCryptoGatewayAdapter(cryptoGateway)
	processors = append(processors, cryptoAdapter)

	// Клиентский код обрабатывает все платежи одинаково
	fmt.Printf("Клиент пытается обработать платеж на $%.2f через %d провайдеров:\n\n", paymentAmount, len(processors))

	for i, processor := range processors {
		fmt.Printf("--- Провайдер #%d ---\n", i+1)
		processor.ProcessPayment(paymentAmount)
		fmt.Println(strings.Repeat("-", 25))
	}
}

// Вспомогательная функция (относится к клиентскому коду, поэтому живет в main)
func printBeverageDetails(b beverage.Beverage) {
	fmt.Printf("Заказ: %s\n", b.GetDescription())
	fmt.Printf("Стоимость: $%.2f\n\n", b.Cost())
}
