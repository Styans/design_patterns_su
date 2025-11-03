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
//   ЧАСТЬ 1: ДЕМОНСТРАЦИЯ "ДЕКОРАТОР"
//
// ==================================================================
func main() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("   ЧАСТЬ 1: ПАТТЕРН \"ДЕКОРАТОР\" (Кафе)")
	fmt.Println(strings.Repeat("=", 60))

	// Задание 3 и 4: Тестирование разных комбинаций

	// Сценарий 1: Кофе (из C# примера)
	fmt.Println("--- Сценарий 1: Базовый Кофе ---")
	var bev1 beverage.IBeverage = &beverage.Coffee{}
	printOrder(bev1)

	// Сценарий 2: Кофе с молоком и сахаром
	fmt.Println("--- Сценарий 2: Кофе с молоком и сахаром ---")
	var bev2 beverage.IBeverage = &beverage.Coffee{}
	bev2 = beverage.NewMilkDecorator(bev2)   // Обернули в молоко
	bev2 = beverage.NewSugarDecorator(bev2)  // Обернули в сахар
	printOrder(bev2)

	// Сценарий 3: Кофе со всем (включая новые добавки)
	fmt.Println("--- Сценарий 3: Кофе с Шоколадом, Ванилью и Молоком ---")
	var bev3 beverage.IBeverage = &beverage.Coffee{}
	bev3 = beverage.NewMilkDecorator(bev3)
	bev3 = beverage.NewChocolateDecorator(bev3) // Новая добавка
	bev3 = beverage.NewVanillaDecorator(bev3)   // Новая добавка
	printOrder(bev3)

	// Сценарий 4: Чай с двойной корицей и сахаром
	fmt.Println("--- Сценарий 4: Чай с двойной Корицей и Сахаром ---")
	var bev4 beverage.IBeverage = &beverage.Tea{}
	bev4 = beverage.NewSugarDecorator(bev4)
	bev4 = beverage.NewCinnamonDecorator(bev4) // Корица 1
	bev4 = beverage.NewCinnamonDecorator(bev4) // Корица 2
	printOrder(bev4)

	// ==================================================================
	//
	//   ЧАСТЬ 2: ДЕМОНСТРАЦИЯ "АДАПТЕР"
	//
	// ==================================================================
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("   ЧАСТЬ 2: ПАТТЕРН \"АДАПТЕР\" (Платежи)")
	fmt.Println(strings.Repeat("=", 60))

	// Задание 3 и 4: Тестирование логики выбора и разных сценариев

	// Сценарий 1: Оплата из США (должна выбраться System A)
	// Клиентский код не знает о 'ExternalPaymentSystemA', он просит процессор у фабрики
	var processor1 payment.IPaymentProcessor = payment.GetPaymentProcessor("US", "USD")
	processor1.ProcessPayment(200.0)
	processor1.RefundPayment(100.0)

	// Сценарий 2: Оплата из Европы (должна выбраться System B)
	var processor2 payment.IPaymentProcessor = payment.GetPaymentProcessor("EU", "EUR")
	processor2.ProcessPayment(300.0)
	processor2.RefundPayment(150.0)

	// Сценарий 3: Оплата из другого региона (должна выбраться внутренняя система)
	var processor3 payment.IPaymentProcessor = payment.GetPaymentProcessor("KZ", "KZT")
	processor3.ProcessPayment(100.0)
	processor3.RefundPayment(50.0)
}

// Вспомогательная функция для вывода заказа
func printOrder(b beverage.IBeverage) {
	fmt.Printf("Описание: %s\n", b.GetDescription())
	fmt.Printf("Стоимость: %.2f\n", b.GetCost())
	fmt.Println(strings.Repeat("-", 30))
}