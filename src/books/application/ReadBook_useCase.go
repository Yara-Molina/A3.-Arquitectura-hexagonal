package application

import (
	"project1/src/books/domain"
	"project1/src/books/domain/entities"
)

type ReadBook struct {
	db domain.IBook
}

func NewReadBook(db domain.IBook) *ReadBook {
	return &ReadBook{db: db}
}

func (rb *ReadBook) ExecuteAll() ([]entities.Book, error) {
	return rb.db.FindAll()
}
