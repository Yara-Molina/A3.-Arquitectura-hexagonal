package routes

import (
	"project1/src/books/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterBookRoutes(router *gin.Engine,
	readController *controllers.ReadBookController,
	createController *controllers.CreateBookController,
	deleteController *controllers.DeleteBookController,
	updateController *controllers.UpdateBookController) {

	router.GET("/books", readController.Handle)
	router.POST("/books", createController.Handle)
	router.DELETE("/books/:id", deleteController.Execute)
	router.PUT("/books/:id", updateController.Handle)
}
