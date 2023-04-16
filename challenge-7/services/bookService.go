package services

import (
	"belajar-gin/entity"
)

type BookService interface {
	AddBook(request entity.BookRequest) (entity.Book, error)
	FindAllBooks() ([]entity.Book, error)
	FindBookById(bookId int) (entity.Book, error)
	UpdateBook(request entity.BookRequest, bookId int) (entity.Book, error)
	DeleteBook(bookId int)
}
