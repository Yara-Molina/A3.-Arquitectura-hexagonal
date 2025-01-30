package domain

import "project1/src/reader/domain/entities"

type IReader interface {
	Safe(reader entities.Reader) (entities.Reader, error)
	LocateAll() ([]entities.Reader, error)
	LocateByID(id int) (entities.Reader, error)
	Upgrade(id int, reader entities.Reader) error
	Erase(id int) error
}
