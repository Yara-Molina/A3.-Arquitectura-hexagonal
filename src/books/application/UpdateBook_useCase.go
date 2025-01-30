package application

import (
	"project1/src/books/domain"
	"project1/src/books/domain/entities"
)

type UpdateBook struct {
	bookRepository domain.IBook
}

func NewUpdateBook(br domain.IBook) *UpdateBook {
	return &UpdateBook{
		bookRepository: br,
	}
}

func (ub *UpdateBook) Update(book entities.Book) error {
	return ub.bookRepository.Update(int(book.ID), book)
}
