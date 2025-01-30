package application

import (
	"project1/src/reader/domain"
	"project1/src/reader/domain/entities"
)

type ViewReader struct {
	db domain.IReader
}

func NewViewReader(db domain.IReader) *ViewReader {
	return &ViewReader{
		db: db,
	}
}

func (rb *ViewReader) ExecuteAll() ([]entities.Reader, error) {
	return rb.db.LocateAll()
}
