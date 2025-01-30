package controllers

import (
	"net/http"
	"project1/src/reader/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EraseReaderController struct {
	useCase *application.Erase
}

func NewEraseReaderController(useCase *application.Erase) *EraseReaderController {
	return &EraseReaderController{useCase: useCase}
}

func (c *EraseReaderController) Execute(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = c.useCase.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Lector eliminado correctamente"})
}
