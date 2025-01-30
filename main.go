package main

import (
	infrastructureB "project1/src/books/infraestructure"
	routesB "project1/src/books/infraestructure/routes"
	infrastructureR "project1/src/reader/infraestructureR"
	routesR "project1/src/reader/infraestructureR/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	readBookController, createBookController, updateBookController, deleteBookController := infrastructureB.Init()

	routesB.RegisterBookRoutes(router, readBookController, createBookController, deleteBookController, updateBookController)

	eraseReaderController, safeReaderController, upgradeReaderController, viewReaderController := infrastructureR.InitReader()

	routesR.RegisterReaderRoutes(router, eraseReaderController, safeReaderController, upgradeReaderController, viewReaderController)

	router.Run(":8080")
}
