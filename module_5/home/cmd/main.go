package main

import (
	"fmt"
	"home/internal/builder"
	"home/internal/prototype"
	"home/internal/singleton"
)

func main() {
	fmt.Println("=== Singleton ===")
	cfg1 := singleton.GetInstance()
	cfg2 := singleton.GetInstance()
	cfg1.Set("AppName", "LibrarySystem")
	fmt.Println(cfg1.Get("AppName"))
	fmt.Println(cfg2.Get("AppName"))
	fmt.Println("Один и тот же объект?", cfg1 == cfg2)

	fmt.Println("\n=== Builder ===")
	director := builder.ReportDirector{}
	textBuilder := &builder.TextReportBuilder{}
	htmlBuilder := &builder.HtmlReportBuilder{}

	textReport := director.ConstructReport(textBuilder, "Отчет", "Содержимое отчета", "Конец")
	htmlReport := director.ConstructReport(htmlBuilder, "HTML Отчет", "Данные", "Низ")

	fmt.Println("\nText Report:\n", textReport)
	fmt.Println("\nHTML Report:\n", htmlReport)

	fmt.Println("\n=== Prototype ===")
	order1 := prototype.Order{
		Products: []prototype.Product{
			{Name: "Телефон", Price: 120000, Quantity: 1},
			{Name: "Наушники", Price: 30000, Quantity: 2},
		},
		Delivery: 2000,
		Discount: prototype.Discount{Description: "Скидка постоянного клиента", Percent: 10},
		Payment:  "Карта",
	}

	order2 := order1.Clone().(*prototype.Order)
	order2.Products[0].Name = "Ноутбук"
	order2.Products[0].Price = 450000

	fmt.Println("Оригинал:", order1.Products[0].Name, order1.Products[0].Price)
	fmt.Println("Клон:", order2.Products[0].Name, order2.Products[0].Price)
}
