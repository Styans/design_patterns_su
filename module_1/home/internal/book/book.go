package book

type Book struct {
	Title     string
	Author    string
	ISBN      string
	Copies    int
	Borrowed  int
}

func NewBook(title, author, isbn string, copies int) *Book {
	return &Book{
		Title:  title,
		Author: author,
		ISBN:   isbn,
		Copies: copies,
	}
}

func (b *Book) IsAvailable() bool {
	return b.Borrowed < b.Copies
}

func (b *Book) Borrow() bool {
	if b.IsAvailable() {
		b.Borrowed++
		return true
	}
	return false
}

func (b *Book) Return() bool {
	if b.Borrowed > 0 {
		b.Borrowed--
		return true
	}
	return false
}
