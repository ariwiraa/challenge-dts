package services

import (
	"gin-gorm/entity"
	"gin-gorm/repository"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
}

func NewBookServiceImpl(bookRepository repository.BookRepository) BookService {
	return &BookServiceImpl{
		BookRepository: bookRepository,
	}
}

// AddBook implements BookService
func (service *BookServiceImpl) AddBook(request entity.BookRequest) (entity.Book, error) {

	// binding ke struct book
	book := entity.Book{
		Title:  request.Title,
		Author: request.Author,
		Desc:   request.Desc,
	}

	book, err := service.BookRepository.AddBook(book)
	if err != nil {
		return book, err
	}

	return book, nil
}

// DeleteBook implements BookService
func (service *BookServiceImpl) DeleteBook(book entity.Book) {

	book, err := service.BookRepository.FindBookById(book.Id)
	if err != nil {
		panic(err)
	}

	service.BookRepository.DeleteBook(book)
}

// GetAllBooks implements BookService
func (service *BookServiceImpl) FindAllBooks() ([]entity.Book, error) {

	books, err := service.BookRepository.FindAllBooks()
	if err != nil {
		return books, err
	}

	return books, nil
}

// GetBookById implements BookService
func (service *BookServiceImpl) FindBookById(bookId int) (entity.Book, error) {

	book, err := service.BookRepository.FindBookById(bookId)
	if err != nil {
		panic(err)
	}

	return book, nil

}

// UpdateBook implements BookService
func (service *BookServiceImpl) UpdateBook(request entity.BookRequest, bookId int) (entity.Book, error) {

	book, err := service.BookRepository.FindBookById(bookId)
	if err != nil {
		panic(err)
	}

	book.Title = request.Title
	book.Author = request.Author
	book.Desc = request.Desc

	book, err = service.BookRepository.UpdateBook(book, bookId)
	if err != nil {
		panic(err)
	}

	return book, nil

}
