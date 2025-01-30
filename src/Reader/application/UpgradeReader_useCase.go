package application

import (
	"project1/src/reader/domain"
	"project1/src/reader/domain/entities"
)

type UpgradeReader struct {
	db domain.IReader
}

func NewUpgradeReader(br domain.IReader) *UpgradeReader {
	return &UpgradeReader{
		db: br,
	}
}

func (ub *UpgradeReader) Upgrade(reader entities.Reader) error {
	return ub.db.Upgrade(int(reader.ID), reader)
}
