package main

import (
	"fmt"
	"strings"

	// Импортируем наши внутренние пакеты
	"shop_system/internal/logistics"
	"shop_system/internal/reporting"
)

func main() {
	// --- Демонстрация "Декоратора" (Система Отчетов) ---
	runDecoratorDemo()

	// --- Демонстрация "Адаптера" (Система Логистики) ---
	runAdapterDemo()
}

// --- ЧАСТЬ 1: ПАТТЕРН "ДЕКОРАТОР" ---
func runDecoratorDemo() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("   ЧАСТЬ 1: ПАТТЕРН \"ДЕКОРАТОР\" (Система Отчетов)")
	fmt.Println(strings.Repeat("=", 60))

	// Задание 4: Тестирование разных комбинаций

	// Тест 1: Простой отчет по пользователям
	fmt.Println("--- Тест 1: Простой Отчет по Пользователям ---")
	simpleUser := &reporting.UserReport{}
	fmt.Println(simpleUser.Generate())

	// Тест 2: Отчет по продажам с сортировкой и PDF экспортом
	fmt.Println("--- Тест 2: Отчет по Продажам (Сортировка, PDF) ---")
	var salesPDF reporting.IReport = &reporting.SalesReport{}
	salesPDF = reporting.NewSortingDecorator(salesPDF, "Date")
	salesPDF = reporting.NewPdfExportDecorator(salesPDF)
	fmt.Println(salesPDF.Generate())

	// Тест 3: Сложный отчет (Фильтры + Сортировка + CSV)
	fmt.Println("--- Тест 3: Сложный Отчет (Фильтры, Сортировка, CSV) ---")
	var complexReport reporting.IReport = &reporting.SalesReport{}
	complexReport = reporting.NewDateFilterDecorator(complexReport, "2025-11-01", "2025-11-02")
	// Задание 2: Новый декоратор
	complexReport = reporting.NewAmountFilterDecorator(complexReport, 90.0)
	complexReport = reporting.NewSortingDecorator(complexReport, "Amount")
	complexReport = reporting.NewCsvExportDecorator(complexReport)
	fmt.Println(complexReport.Generate())

	// Задание 3: Тестирование динамического конструктора
	fmt.Println("--- Тест 4: Динамический Конструктор (Задание 3) ---")
	request := reporting.ReportRequest{
		ReportType: "user",
		SortBy:     "Region",
		ExportAs:   "csv",
	}
	dynamicReport := reporting.BuildReport(request)
	if dynamicReport != nil {
		fmt.Println(dynamicReport.Generate())
	}
}

// --- ЧАСТЬ 2: ПАТТЕРН "АДАПТЕР" ---
func runAdapterDemo() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("   ЧАСТЬ 2: ПАТТЕРН \"АДАПТЕР\" (Система Логистики)")
	fmt.Println(strings.Repeat("=", 60))

	// Задания 2, 3, 4, 5: Тестируем всех провайдеров через фабрику

	// Провайдеры для теста
	providers := []string{
		"internal", // Внутренняя служба
		"serviceA", // Адаптер А
		"serviceB", // Адаптер Б
		"serviceC", // Задание 2: Адаптер C
		"serviceD", // Несуществующий провайдер для теста ошибок
	}
	
	orderId := "451" // Фиктивный ID заказа

	for _, provider := range providers {
		fmt.Printf("\n--- Тест: Провайдер '%s' | Заказ '%s' ---\n", provider, orderId)

		// Получаем сервис из Фабрики (Задание 4)
		service, err := logistics.GetDeliveryService(provider)
		if err != nil {
			// Задание 4: Обработка ошибок
			fmt.Printf("Ошибка Фабрики: %v\n", err)
			continue
		}

		// Задание 5: Расчет стоимости
		cost := service.CalculateDeliveryCost(orderId)
		fmt.Printf("Расчетная стоимость: %.2f\n", cost)

		// Доставка
		deliveryStatus := service.DeliverOrder(orderId)
		fmt.Printf("Статус отправки: %s\n", deliveryStatus)

		// Проверка статуса
		status := service.GetDeliveryStatus(orderId)
		fmt.Printf("Текущий статус: %s\n", status)
	}
}
