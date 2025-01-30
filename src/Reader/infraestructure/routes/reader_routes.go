package routes

import (
	"project1/src/reader/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterReaderRoutes(router *gin.Engine,
	viewController *controllers.ViewReaderController,
	safeController *controllers.SafeReaderController,
	upgradeController *controllers.UpgradeReaderController,
	eraseController *controllers.EraseReaderController) {

	router.GET("/readers", viewController.Handle)
	router.POST("/readers", safeController.Handle)
	router.DELETE("/readers/:id", eraseController.Execute)
	router.PUT("/readers/:id", upgradeController.Handle)
}
