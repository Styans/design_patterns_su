package main

import (
	"home/internal/dry"
	"home/internal/kiss"
	"home/internal/yagni"
)

func main() {
	logger := dry.Logger{}
	logger.Log("ERROR", "Ошибка подключения")
	logger.Log("INFO", "Приложение запущено")

	config := dry.Config{ConnectionString: "Server=myServer;Database=myDb;User Id=myUser;Password=myPass;"}
	db := dry.DatabaseService{Config: config}
	db.Connect()

	logService := dry.LoggingService{Config: config}
	logService.Log("Test message")

	kiss.ProcessNumbers([]int{1, -2, 3, 0, 5})
	kiss.PrintPositiveNumbers([]int{1, -1, 2, 3})
	kiss.Divide(10, 2)
	kiss.Divide(10, 0)

	user := yagni.User{Name: "Алихан", Email: "test@mail.com"}
	user.SaveToDatabase()

	reader := yagni.FileReader{}
	reader.ReadFile("file.txt")

	report := yagni.ReportGenerator{}
	report.GeneratePdfReport()
}
