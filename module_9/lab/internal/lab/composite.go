package lab

import "fmt"

// FileSystemComponent – общий интерфейс для файлов и директорий.
type FileSystemComponent interface {
	Display(depth int)
}

// ----------- Лист: File -----------

type File struct {
	name string
}

func NewFile(name string) *File {
	return &File{name: name}
}

func (f *File) Display(depth int) {
	fmt.Printf("%s File: %s\n", dash(depth), f.name)
}

// ----------- Контейнер: Directory -----------

type Directory struct {
	name     string
	children []FileSystemComponent
}

func NewDirectory(name string) *Directory {
	return &Directory{name: name, children: make([]FileSystemComponent, 0)}
}

func (d *Directory) Add(component FileSystemComponent) {
	d.children = append(d.children, component)
}

func (d *Directory) Remove(component FileSystemComponent) {
	for i, c := range d.children {
		if c == component {
			d.children = append(d.children[:i], d.children[i+1:]...)
			break
		}
	}
}

func (d *Directory) GetChild(index int) FileSystemComponent {
	if index < 0 || index >= len(d.children) {
		return nil
	}
	return d.children[index]
}

func (d *Directory) Display(depth int) {
	fmt.Printf("%s Directory: %s\n", dash(depth), d.name)
	for _, c := range d.children {
		c.Display(depth + 2)
	}
}

// вспомогательная функция для отступов
func dash(depth int) string {
	return string(make([]rune, depth, depth))
}
