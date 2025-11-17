package main

import (
	"fmt"

	"home/internal/facade"
	"home/internal/filesystem"
)

func main() {
	runFacadeDemo()
	fmt.Println()
	runCompositeDemo()
}

func runFacadeDemo() {
	ht := facade.NewHomeTheaterFacade()

	ht.WatchMovie("Inception")
	ht.SetVolume(70)
	ht.StopMovie()

	fmt.Println()

	ht.PlayGame("The Witcher 3")
	fmt.Println()
	ht.ListenMusic()
}

func runCompositeDemo() {
	// Корневая папка
	root := filesystem.NewDirectory("root")

	// Папка docs
	docs := filesystem.NewDirectory("docs")
	docs.Add(filesystem.NewFile("report.docx", 42_000))
	docs.Add(filesystem.NewFile("slides.pptx", 80_000))

	// Папка images
	images := filesystem.NewDirectory("images")
	images.Add(filesystem.NewFile("photo1.jpg", 1_500_000))
	images.Add(filesystem.NewFile("photo2.png", 2_100_000))

	// Вложенная папка внутри images
	sub := filesystem.NewDirectory("icons")
	sub.Add(filesystem.NewFile("icon.png", 12_000))
	images.Add(sub)

	// Собираем структуру
	root.Add(docs)
	root.Add(images)
	root.Add(filesystem.NewFile("readme.txt", 1_000))

	// Попытка добавить дубликат
	root.Add(filesystem.NewFile("readme.txt", 2_000))

	fmt.Println("=== File system structure ===")
	root.Display("")

	fmt.Printf("\nTotal size of %q: %d bytes\n", root.Name(), root.Size())

	// Пример удаления
	root.Remove("readme.txt")
	root.Remove("no-such-file")

	fmt.Println("\n=== After remove ===")
	root.Display("")
	fmt.Printf("\nTotal size of %q: %d bytes\n", root.Name(), root.Size())
}
