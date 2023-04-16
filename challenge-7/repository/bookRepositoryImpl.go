package repository

import (
	"belajar-gin/entity"
	"database/sql"
	"errors"
)

type BookRepositoryImpl struct {
	Db *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &BookRepositoryImpl{Db: db}
}

func (repository *BookRepositoryImpl) AddBook(book entity.Book) (entity.Book, error) {

	query := "INSERT INTO books(title, author, description) VALUES($1, $2, $3) RETURNING *"

	err := repository.Db.QueryRow(query, book.Title, book.Author, book.Desc).Scan(&book.Id, &book.Title, &book.Author, &book.Desc)
	if err != nil {
		panic(err)
	}
	return book, nil
}

// DeleteBook implements BookRepository
func (repository *BookRepositoryImpl) DeleteBook(bookId int) {
	query := "DELETE FROM books where id = $1"
	_, err := repository.Db.Exec(query, bookId)
	if err != nil {
		panic(err)
	}
}

// FindAllBooks implements BookRepository
func (repository *BookRepositoryImpl) FindAllBooks() ([]entity.Book, error) {
	query := "SELECT * FROM books ORDER BY id DESC"
	rows, err := repository.Db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var books []entity.Book
	for rows.Next() {
		book := entity.Book{}
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Desc)
		if err != nil {
			return books, errors.New("tidak dapat mengambil semua buku")
		}
		books = append(books, book)
	}
	return books, nil
}

// FindBookById implements BookRepository
func (repository *BookRepositoryImpl) FindBookById(bookId int) (entity.Book, error) {
	query := "SELECT * FROM books WHERE id = $1"
	rows, err := repository.Db.Query(query, bookId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	book := entity.Book{}

	if rows.Next() {
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Desc)
		if err != nil {
			panic(err)
		}
		return book, nil
	} else {
		return book, errors.New("buku tidak ditemukan")
	}
}

// UpdateBook implements BookRepository
func (repository *BookRepositoryImpl) UpdateBook(book entity.Book, bookId int) (entity.Book, error) {
	query := "UPDATE books SET title = $1, author = $2, description = $3 WHERE id = $4"
	_, err := repository.Db.Exec(query, book.Title, book.Author, book.Desc, bookId)
	if err != nil {
		panic(err)
	}

	return book, nil
}
