package services

import (
	"gin-gorm/entity"
)

type BookService interface {
	AddBook(request entity.BookRequest) (entity.Book, error)
	FindAllBooks() ([]entity.Book, error)
	FindBookById(bookId int) (entity.Book, error)
	UpdateBook(request entity.BookRequest, bookId int) (entity.Book, error)
	DeleteBook(book entity.Book)
}
