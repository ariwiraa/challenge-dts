package repository

import (
	"errors"
	"gin-gorm/entity"
	"log"

	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
	Db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &BookRepositoryImpl{Db: db}
}

func (repository *BookRepositoryImpl) AddBook(book entity.Book) (entity.Book, error) {

	err := repository.Db.Create(&book).Error
	if err != nil {
		log.Fatal("error creating book data: ", err)
	}
	return book, nil
}

// DeleteBook implements BookRepository
func (repository *BookRepositoryImpl) DeleteBook(book entity.Book) {
	err := repository.Db.Where("id = ?", book.Id).Delete(&book).Error
	if err != nil {
		log.Fatal("error deleting data: ", err.Error())
		return
	}
}

// FindAllBooks implements BookRepository
func (repository *BookRepositoryImpl) FindAllBooks() ([]entity.Book, error) {
	var books []entity.Book

	err := repository.Db.Find(&books).Error
	if err != nil {
		log.Fatal("error getting all books data: ", err)
	}
	return books, nil
}

// FindBookById implements BookRepository
func (repository *BookRepositoryImpl) FindBookById(bookId int) (entity.Book, error) {
	var book entity.Book

	err := repository.Db.First(&book, "id = ?", bookId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Fatal("book is not found")
		}
		log.Fatal("error getting book data: ", err)
	}
	return book, nil
}

// UpdateBook implements BookRepository
func (repository *BookRepositoryImpl) UpdateBook(book entity.Book, bookId int) (entity.Book, error) {

	err := repository.Db.Model(&book).Where("id = ?", bookId).Updates(&book).Error
	if err != nil {
		log.Fatal("error updating data: ", err)
	}

	return book, nil
}
