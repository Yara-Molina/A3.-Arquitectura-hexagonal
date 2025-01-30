package application

import (
	"project1/src/books/domain"
	"project1/src/books/domain/entities"
)

type CreateBook struct {
	db domain.IBook
}

func NewCreateBook(db domain.IBook) *CreateBook {
	return &CreateBook{db: db}
}

func (uc *CreateBook) Execute(book entities.Book) (entities.Book, error) {
	return uc.db.Save(book)
}
