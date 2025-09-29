package main

import (
	"fmt"
	"practice/internal/creater"
	"strings"
)

func main() {

	var choice string
	fmt.Println("Выберете тип документа: report / resume / letter")
	fmt.Scan(&choice)

	choice = strings.ToLower(choice)
	c := creater.GetCreator(choice)
	var data string
	switch choice {
	case "report":
		fmt.Println("Введите заголовок отчета:")
	case "resume":
		fmt.Println("Введите имя владельца резюме:")
	case "letter":
		fmt.Println("Введите содержание письма:")
	default:
		fmt.Println("Ошибка: неизвестный тип документа")
		return
	}
	fmt.Scan(&data)
	doc := c.CreateDocument(data)
	doc.Open()
}
