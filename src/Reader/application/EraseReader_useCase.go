package application

import (
	"fmt"
	"project1/src/reader/domain"
)

type Erase struct {
	db domain.IReader
}

func NewEraseReader(db domain.IReader) *Erase {
	return &Erase{db: db}
}

func (uc *Erase) Execute(readerID int) error {
	fmt.Println("Lector a eliminar con ID:", readerID)

	err := uc.db.Erase(readerID)
	if err != nil {
		fmt.Println("Error al eliminar lector", err)
		return err
	}

	fmt.Println("Libro eliminado")
	return nil
}
