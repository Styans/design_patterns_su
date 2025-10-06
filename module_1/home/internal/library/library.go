package library

import (
	"fmt"
	"home/internal/book"
	"home/internal/reader"
)

type Library struct {
	Books   []*book.Book
	Readers []*reader.Reader
}

func (l *Library) AddBook(b *book.Book) {
	l.Books = append(l.Books, b)
	fmt.Printf("Книга '%s' добавлена в библиотеку.\n", b.Title)
}

func (l *Library) RemoveBook(isbn string) {
	for i, b := range l.Books {
		if b.ISBN == isbn {
			l.Books = append(l.Books[:i], l.Books[i+1:]...)
			fmt.Printf("Книга '%s' удалена.\n", b.Title)
			return
		}
	}
	fmt.Println("Книга не найдена.")
}

func (l *Library) RegisterReader(r *reader.Reader) {
	l.Readers = append(l.Readers, r)
	fmt.Printf("Читатель '%s' зарегистрирован.\n", r.Name)
}

func (l *Library) RemoveReader(id int) {
	for i, r := range l.Readers {
		if r.ID == id {
			l.Readers = append(l.Readers[:i], l.Readers[i+1:]...)
			fmt.Printf("Читатель '%s' удалён.\n", r.Name)
			return
		}
	}
	fmt.Println("Читатель не найден.")
}

func (l *Library) BorrowBook(isbn string, readerID int) {
	for _, b := range l.Books {
		if b.ISBN == isbn {
			if b.Borrow() {
				fmt.Printf("Книга '%s' выдана читателю #%d.\n", b.Title, readerID)
				return
			}
			fmt.Println("Нет доступных экземпляров книги.")
			return
		}
	}
	fmt.Println("Книга не найдена.")
}

func (l *Library) ReturnBook(isbn string, readerID int) {
	for _, b := range l.Books {
		if b.ISBN == isbn {
			if b.Return() {
				fmt.Printf("Читатель #%d вернул книгу '%s'.\n", readerID, b.Title)
				return
			}
			fmt.Println("Ошибка возврата.")
			return
		}
	}
	fmt.Println("Книга не найдена.")
}

func (l *Library) ListBooks() {
	fmt.Println("\n📚 Список книг:")
	for _, b := range l.Books {
		fmt.Printf("• %s (%s) [%d/%d доступно]\n", b.Title, b.Author, b.Copies-b.Borrowed, b.Copies)
	}
}
