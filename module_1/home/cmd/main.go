package main

import (
	"home/internal/book"
	"home/internal/library"
	"home/internal/reader"
)

func main() {
	lib := &library.Library{}

	book1 := book.NewBook("Война и мир", "Л. Толстой", "111-222", 3)
	book2 := book.NewBook("Преступление и наказание", "Ф. Достоевский", "333-444", 2)

	lib.AddBook(book1)
	lib.AddBook(book2)

	reader1 := reader.NewReader(1, "Алихан")
	reader2 := reader.NewReader(2, "Мерей")

	lib.RegisterReader(reader1)
	lib.RegisterReader(reader2)

	lib.BorrowBook("111-222", 1)
	lib.BorrowBook("111-222", 2)
	lib.BorrowBook("111-222", 2)
	lib.ListBooks()

	lib.ReturnBook("111-222", 1)

	lib.RemoveBook("333-444")

	lib.RemoveReader(1)

	lib.ListBooks()
}
