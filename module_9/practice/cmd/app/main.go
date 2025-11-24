package main

import (
	"fmt"

	"practice/internal/practice"
)

func main() {
	fmt.Println("========== FACADE EXAMPLE (HOTEL) ==========")
	facade := practice.NewHotelFacade()

	// 1) Бронирование номера с рестораном и уборкой
	roomID, tableID := facade.BookRoomWithRestaurantAndCleaning(
		"Jane Doe",
		"Deluxe",
		3,
		101,
		"19:00",
	)
	fmt.Println()

	// 2) Организация мероприятия с оборудованием и номерами
	eventID, groupRoomID := facade.OrganizeEventWithRoomsAndEquipment(
		"Tech Conference",
		"2025-12-01",
		120,
	)
	fmt.Println()
	
	_ = groupRoomID

	// 3) Бронирование стола в ресторане с вызовом такси
	tableID2 := facade.BookRestaurantTableWithTaxi(
		"John Smith",
		4,
		"20:30",
	)
	_ = tableID2

	fmt.Println()

	// Дополнительно: отмена бронирований и уборка по запросу
	facade.CancelRoomBooking(roomID)
	facade.CancelRestaurantBooking(tableID)
	facade.CancelEventBooking(eventID)
	facade.RequestCleaningOnDemand(101)

	fmt.Println()
	fmt.Println("========== COMPOSITE EXAMPLE (ORGANIZATION) ==========")

	// Корневой департамент
	root := practice.NewDepartment("Head Office")

	// Отделы
	devDept := practice.NewDepartment("Development")
	hrDept := practice.NewDepartment("HR")

	// Сотрудники
	dev1 := practice.NewEmployee("Alice", "Backend Developer", 500000)
	dev2 := practice.NewEmployee("Bob", "Frontend Developer", 480000)
	hr1 := practice.NewEmployee("Eve", "HR Manager", 450000)

	// Контрактор
	contractor := practice.NewContractor("Charlie", "DevOps Contractor", 300000)

	// Формируем структуру
	devDept.Add(dev1)
	devDept.Add(dev2)
	devDept.Add(contractor)

	hrDept.Add(hr1)

	root.Add(devDept)
	root.Add(hrDept)

	// Отображение структуры
	root.Display(1)

	// Общий бюджет и штат
	fmt.Printf("\nTotal budget of %s = %.2f\n", root.GetName(), root.GetBudget())
	fmt.Printf("Total headcount of %s = %d\n", root.GetName(), root.GetHeadcount())

	// Изменение зарплаты и пересчёт бюджета
	dev1.SetSalary(550000)
	fmt.Printf("\nAfter salary change of %s:\n", dev1.GetName())
	fmt.Printf("Total budget of %s = %.2f\n", root.GetName(), root.GetBudget())

	// Поиск сотрудника по имени
	fmt.Println()
	nameToFind := "Alice"
	found := root.FindByName(nameToFind)
	if found != nil {
		fmt.Printf("Found component with name %s:\n", nameToFind)
		found.Display(2)
	} else {
		fmt.Printf("Component with name %s not found\n", nameToFind)
	}

	// Список всех сотрудников и контракторов отдела
	fmt.Println()
	root.ListAllEmployees()
}
