package routes

import (
	"project1/src/reader/infraestructureR/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterReaderRoutes(router *gin.Engine,
	viewReaderController *controllers.ViewReaderController,
	safeReaderController *controllers.SafeReaderController,
	upgradeReaderController *controllers.UpgradeReaderController,
	eraseReaderController *controllers.EraseReaderController) {

	router.GET("/readers", viewReaderController.Handle)
	router.POST("/readers", safeReaderController.Handle)
	router.PUT("/readers/:id", upgradeReaderController.Handle)
	router.DELETE("/readers/:id", eraseReaderController.Execute)
}
