package controllers

import (
	"project1/src/reader/application"
	"project1/src/reader/domain/entities"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpgradeReaderController struct {
	useCase *application.UpgradeReader
}

func NewUpgradeReaderController(uc *application.UpgradeReader) *UpgradeReaderController {
	return &UpgradeReaderController{useCase: uc}
}

func (c *UpgradeReaderController) Handle(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var reader entities.Reader
	if err := ctx.ShouldBindJSON(&reader); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reader.ID = int32(id)

	err = c.useCase.Upgrade(reader)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Libro actualizado correctamente"})
}
