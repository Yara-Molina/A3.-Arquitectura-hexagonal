package application

import (
	"project1/src/reader/domain"
	"project1/src/reader/domain/entities"
)

type SafeReader struct {
	db domain.IReader
}

func NewSafeReader(db domain.IReader) *SafeReader {
	return &SafeReader{db: db}
}

func (uc *SafeReader) Execute(reader entities.Reader) (entities.Reader, error) {
	return uc.db.Safe(reader)
}
