package repository

import (
	"belajar-gin/entity"
)

type BookRepository interface {
	AddBook(book entity.Book) (entity.Book, error)
	UpdateBook(book entity.Book, bookId int) (entity.Book, error)
	DeleteBook(bookId int)
	FindBookById(bookId int) (entity.Book, error)
	FindAllBooks() ([]entity.Book, error)
}
