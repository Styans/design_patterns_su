package document

import "fmt"

type IDocument interface {
	Open()
}

type Report struct {
	Title string
}

func (r *Report) Open() {
	fmt.Println("Opening report:", r.Title)
}

type Resume struct {
	Owner string
}

func (r *Resume) Open() {
	fmt.Println("Opening resume of:", r.Owner)
}

type Letter struct {
	Content string
}

func (l *Letter) Open() {
	fmt.Println("Opening letter with content:", l.Content)
}
