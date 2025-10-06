package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"practice/internal/builder"
	"practice/internal/logger"
	"practice/internal/prototype"
)

func demoLogger() {
	cfg := `{"level":"INFO","file_path":"application.log"}`
	_ = osWrite("logger_config.json", []byte(cfg))
	log := logger.GetInstance()
	_ = log.LoadConfig("logger_config.json")
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				level := logger.INFO
				switch rand.Intn(3) {
				case 0:
					level = logger.INFO
				case 1:
					level = logger.WARNING
				case 2:
					level = logger.ERROR
				}
				_ = log.Log(fmt.Sprintf("goroutine %d message %d", id, j), level)
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			}
		}(i)
	}
	wg.Wait()
	log.SetLogLevel(logger.ERROR)
	_ = log.Log("this is an info after level change", logger.INFO)
	_ = log.Log("this is an error after level change", logger.ERROR)
	log.Close()
	reader := logger.NewLogReader("application.log")
	entries, _ := reader.ReadByLevel(logger.ERROR)
	fmt.Println("Errors found:", len(entries))
	for _, e := range entries {
		fmt.Println(e.Timestamp.Format(time.RFC3339), e.Level.String(), e.Message)
	}
}

func osWrite(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}

func demoBuilder() {
	dir := builder.ReportDirector{}
	text := &builder.TextReportBuilder{}
	html := &builder.HtmlReportBuilder{}
	pdf := &builder.PdfReportBuilder{}
	sections := []builder.Section{
		{Title: "Intro", Content: "Introduction content"},
		{Title: "Data", Content: "Data content"},
	}
	style := builder.ReportStyle{BgColor: "#ffffff", TextColor: "#000000", FontSize: 14}
	rep1 := dir.Construct(text, "Report 1", "Main content", "Footer", sections, style)
	rep2 := dir.Construct(html, "Report 2", "HTML content", "Footer HTML", sections, style)
	rep3 := dir.Construct(pdf, "Report 3", "PDF content", "Footer PDF", sections, style)
	_ = rep1.ExportText("report1.txt")
	_ = rep2.ExportHTML("report2.html")
	_ = rep3.ExportPDF("report3.pdf")
	fmt.Println("Reports exported")
}

func demoPrototype() {
	base := prototype.Character{
		Name:      "Warrior",
		Health:    200,
		Strength:  30,
		Agility:   10,
		Intellect: 5,
		Weapon:    prototype.Weapon{Name: "Sword", Damage: 25},
		Armor:     prototype.Armor{Name: "Plate", Defense: 15},
		Skills: []prototype.Skill{
			{Name: "Slash", Power: 10},
			{Name: "Roar", Power: 5},
		},
	}
	clone := base.Clone()
	clone.Name = "Warrior Clone"
	clone.Weapon.Name = "Greatsword"
	clone.Skills[0].Power = 20
	fmt.Println("Base Weapon:", base.Weapon.Name, "Clone Weapon:", clone.Weapon.Name)
	fmt.Println("Base Skill Power:", base.Skills[0].Power, "Clone Skill Power:", clone.Skills[0].Power)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Demo Singleton Logger")
	demoLogger()
	fmt.Println("\nDemo Builder")
	demoBuilder()
	fmt.Println("\nDemo Prototype")
	demoPrototype()
}
