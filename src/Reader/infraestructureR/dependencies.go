package infrastructure

import (
	"project1/src/reader/application"
	"project1/src/reader/infraestructureR/controllers"
)

func InitReader() (*controllers.ViewReaderController, *controllers.SafeReaderController, *controllers.UpgradeReaderController, *controllers.EraseReaderController) {
	readerRepository := NewMySQL()

	viewReaderUseCase := application.NewViewReader(readerRepository)
	safeReaderUseCase := application.NewSafeReader(readerRepository)
	upgradeReaderUseCase := application.NewUpgradeReader(readerRepository)
	eraseReaderUseCase := application.NewEraseReader(readerRepository)

	viewReaderController := controllers.NewViewReaderController(viewReaderUseCase)
	safeReaderController := controllers.NewSafeReaderController(safeReaderUseCase)
	upgradeReaderController := controllers.NewUpgradeReaderController(upgradeReaderUseCase)
	eraseReaderController := controllers.NewEraseReaderController(eraseReaderUseCase)

	return viewReaderController, safeReaderController, upgradeReaderController, eraseReaderController
}
