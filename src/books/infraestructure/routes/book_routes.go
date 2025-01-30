package routes

import (
	"project1/src/books/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterBookRoutes(router *gin.Engine,
	readBookController *controllers.ReadBookController,
	createBookController *controllers.CreateBookController,
	deleteBookController *controllers.DeleteBookController,
	updateBookController *controllers.UpdateBookController) {

	router.GET("/books", readBookController.Handle)
	router.POST("/books", createBookController.Handle)
	router.PUT("/books/:id", updateBookController.Handle)
	router.DELETE("/books/:id", deleteBookController.Execute)
}
