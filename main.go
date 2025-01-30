package main

import (
	infrastructure "project1/src/books/infraestructure"
	"project1/src/books/infraestructure/routes"
	infrastructure "project1/src/reader/infraestructure"
	"project1/src/reader/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	readBookController, createBookController, updateBookController, deleteBookController := infrastructure.Init()

	viewReaderController, safeReaderController, upgradeReaderController, eraseReaderController := infrastructure.Init()

	routes.RegisterBookRoutes(router, readBookController, createBookController, deleteBookController, updateBookController)

	readerRoutes.RegisterReaderRoutes(router, viewReaderController, safeReaderController, upgradeReaderController, eraseReaderController)

	router.Run(":8080")
}
