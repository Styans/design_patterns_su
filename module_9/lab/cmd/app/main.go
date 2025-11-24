// cmd/app/main.go
package main

import (
	"fmt"

	"lab/internal/lab"
)

func main() {
	// ===== Демонстрация Facade =====
	fmt.Println("=== FACADE DEMO ===")
	theater := lab.NewHomeTheaterFacade()

	theater.StartMovie()
	fmt.Println()
	theater.EndMovie()

	fmt.Println()

	// ===== Демонстрация Composite =====
	fmt.Println("=== COMPOSITE DEMO ===")

	root := lab.NewDirectory("Root")
	file1 := lab.NewFile("File1.txt")
	file2 := lab.NewFile("File2.txt")

	subDir := lab.NewDirectory("SubDirectory")
	subFile1 := lab.NewFile("SubFile1.txt")

	// Формируем структуру
	root.Add(file1)
	root.Add(file2)
	subDir.Add(subFile1)
	root.Add(subDir)

	// Отображаем
	root.Display(1)
}
