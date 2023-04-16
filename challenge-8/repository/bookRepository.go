package repository

import (
	"gin-gorm/entity"
)

type BookRepository interface {
	AddBook(book entity.Book) (entity.Book, error)
	UpdateBook(book entity.Book, bookId int) (entity.Book, error)
	DeleteBook(book entity.Book)
	FindBookById(bookId int) (entity.Book, error)
	FindAllBooks() ([]entity.Book, error)
}
