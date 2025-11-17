package filesystem

import "fmt"

// FileSystemComponent — общий интерфейс для файла и папки.
type FileSystemComponent interface {
	Name() string
	Display(indent string)
	Size() int64
}

// --- File ---

type File struct {
	name string
	size int64
}

func NewFile(name string, size int64) *File {
	return &File{name: name, size: size}
}

func (f *File) Name() string { return f.name }

func (f *File) Display(indent string) {
	fmt.Printf("%s- %s (%d bytes)\n", indent, f.name, f.size)
}

func (f *File) Size() int64 { return f.size }

// --- Directory ---

type Directory struct {
	name     string
	children []FileSystemComponent
}

func NewDirectory(name string) *Directory {
	return &Directory{name: name}
}

func (d *Directory) Name() string { return d.name }

func (d *Directory) Display(indent string) {
	fmt.Printf("%s+ %s/\n", indent, d.name)
	nextIndent := indent + "  "
	for _, c := range d.children {
		c.Display(nextIndent)
	}
}

func (d *Directory) Size() int64 {
	var total int64
	for _, c := range d.children {
		total += c.Size()
	}
	return total
}

func (d *Directory) findIndexByName(name string) int {
	for i, c := range d.children {
		if c.Name() == name {
			return i
		}
	}
	return -1
}

// Add — добавляет компонент, проверяя, что с таким именем ещё нет.
func (d *Directory) Add(c FileSystemComponent) {
	if d.findIndexByName(c.Name()) != -1 {
		fmt.Printf("Directory %q already contains %q, skip add\n", d.name, c.Name())
		return
	}
	d.children = append(d.children, c)
}

// Remove — удаляет компонент по имени, если он есть.
func (d *Directory) Remove(name string) {
	i := d.findIndexByName(name)
	if i == -1 {
		fmt.Printf("Directory %q: component %q not found, nothing to remove\n", d.name, name)
		return
	}
	d.children = append(d.children[:i], d.children[i+1:]...)
}
