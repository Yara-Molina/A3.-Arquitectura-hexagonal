package application

import (
	"fmt"
	"project1/src/books/domain"
)

type DeleteBook struct {
	db domain.IBook
}

func NewDeleteBook(db domain.IBook) *DeleteBook {
	return &DeleteBook{db: db}
}

func (uc *DeleteBook) Execute(bookID int) error {
	fmt.Println("libro a eliminar con ID:", bookID)

	err := uc.db.Delete(bookID)
	if err != nil {
		fmt.Println("Error al eliminar producto:", err)
		return err
	}

	fmt.Println("Libro eliminado correctamente")
	return nil

}
