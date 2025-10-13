package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"lab/internal/builder"
	"lab/internal/prototype"
	"lab/internal/singleton"
)

//======================================== Демка синглтона ========================================

func demoLogger() {
	cfg := `{"level":"INFO","file_path":"application.log"}`
	_ = os.WriteFile("logger_config.json", []byte(cfg), 0644)
	log := singleton.GetInstance()
	_ = log.LoadConfig("logger_config.json")
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 20; j++ {
				level := singleton.INFO
				switch rand.Intn(3) {
				case 0:
					level = singleton.INFO
				case 1:
					level = singleton.WARNING
				case 2:
					level = singleton.ERROR
				}
				_ = log.Log(fmt.Sprintf("goroutine %d message %d", id, j), level)
				time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
			}
		}(i)
	}
	wg.Wait()
	log.SetLogLevel(singleton.ERROR)
	_ = log.Log("info after level change", singleton.INFO)
	_ = log.Log("error after level change", singleton.ERROR)
	_ = log.Close()
	reader := singleton.NewLogReader("application.log")
	entries, err := reader.ReadByLevel(singleton.ERROR)
	if err == nil {
		fmt.Println("Errors found:", len(entries))
		for _, e := range entries {
			fmt.Println(e.Timestamp.Format(time.RFC3339), e.Level.String(), e.Message)
		}
	}
}

//======================================== Демка билдера ========================================

func demoBuilder() {
	officeBuilder := builder.NewOfficeComputerBuilder()
	director := builder.NewComputerDirector(officeBuilder)
	office := director.ConstructComputer()

	gamingBuilder := builder.NewGamingComputerBuilder()
	director2 := builder.NewComputerDirector(gamingBuilder)
	gaming := director2.ConstructComputer()

	fmt.Println("Office computer")
	fmt.Println(office.String())
	fmt.Println("Gaming computer")
	fmt.Println(gaming.String())
}

// ======================================== Демка прототипа ========================================
func demoPrototype() {
	sec1 := prototype.Section{Title: "Intro", Content: "Welcome", Images: []prototype.Image{{URL: "img1.png", Alt: "img1"}}}
	sec2 := prototype.Section{Title: "Content", Content: "Body", Images: []prototype.Image{{URL: "img2.png", Alt: "img2"}}}
	doc := prototype.Document{Title: "Doc1", Content: "Main", Sections: []prototype.Section{sec1, sec2}}
	manager := prototype.DocumentManager{}
	doc2 := manager.CreateDocument(doc).(prototype.Document)
	doc2.Title = "Doc2"
	doc2.Sections[0].Content = "Changed"
	fmt.Println("Original title:", doc.Title)
	fmt.Println("Clone title:", doc2.Title)
	fmt.Println("Original section content:", doc.Sections[0].Content)
	fmt.Println("Clone section content:", doc2.Sections[0].Content)
}

// ======================================== Main ========================================

func main() {
	fmt.Println("Demo Singleton Logger")
	demoLogger()
	fmt.Println("\nDemo Builder")
	demoBuilder()
	fmt.Println("\nDemo Prototype")
	demoPrototype()
}
