package domain

import "project1/src/books/domain/entities"

type IBook interface {
	Save(book entities.Book) (entities.Book, error)
	FindAll() ([]entities.Book, error)
	FindByID(id int) (entities.Book, error)
	Update(id int, book entities.Book) error
	Delete(id int) error
}
