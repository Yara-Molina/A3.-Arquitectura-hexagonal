package controllers

import (
	"net/http"
	"project1/src/reader/application"
	"project1/src/reader/domain/entities"

	"github.com/gin-gonic/gin"
)

type SafeReaderController struct {
	useCase *application.SafeReader
}

func NewSafeReaderController(uc *application.SafeReader) *SafeReaderController {
	return &SafeReaderController{useCase: uc}

}

func (c *SafeReaderController) Handle(ctx *gin.Context) {
	var newReader entities.Reader

	if err := ctx.ShouldBindJSON(&newReader); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	book, err := c.useCase.Execute(newReader)
	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, book)
}
