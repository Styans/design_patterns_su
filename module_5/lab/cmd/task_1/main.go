package main

import (
	"fmt"
	"lab/internal/singleton"
)

func main() {
	logger := singleton.GetInstance()
	logger.Log("Система запущена", singleton.INFO)
	logger.Log("Возможная ошибка подключения", singleton.WARNING)
	logger.Log("Фатальная ошибка!", singleton.ERROR)

	reader := singleton.NewLogReader("app.log")
	entries, _ := reader.ReadByLevel(singleton.WARNING)
	for _, e := range entries {
		fmt.Println(e.Timestamp, e.Level, e.Message)
	}
}
