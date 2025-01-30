package infrastructure

import (
	"project1/src/books/application"
	"project1/src/books/infraestructure/controllers"
)

func Init() (*controllers.ReadBookController, *controllers.CreateBookController, *controllers.UpdateBookController, *controllers.DeleteBookController) {
	bookRepository := NewMySQL()

	readBookUseCase := application.NewReadBook(bookRepository)
	createBookUseCase := application.NewCreateBook(bookRepository)
	updateBookUseCase := application.NewUpdateBook(bookRepository)
	deleteBookUseCase := application.NewDeleteBook(bookRepository)

	readBookController := controllers.NewReadBookController(readBookUseCase)
	createBookController := controllers.NewCreateBookController(createBookUseCase)
	updateBookController := controllers.NewUpdateBookController(updateBookUseCase)
	deleteBookController := controllers.NewDeleteBookController(deleteBookUseCase)

	return readBookController, createBookController, updateBookController, deleteBookController
}
