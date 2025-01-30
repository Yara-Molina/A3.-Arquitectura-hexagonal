package controllers

import (
	"net/http"
	"project1/src/reader/application"

	"github.com/gin-gonic/gin"
)

type ViewReaderController struct {
	useCase *application.ViewReader
}

func NewViewReaderController(uc *application.ViewReader) *ViewReaderController {
	return &ViewReaderController{useCase: uc}
}

func (c *ViewReaderController) Handle(ctx *gin.Context) {
	reader, err := c.useCase.ExecuteAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, reader)
}
