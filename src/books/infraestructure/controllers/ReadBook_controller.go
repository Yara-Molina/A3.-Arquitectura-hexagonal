package controllers

import (
	"net/http"
	"project1/src/books/application"

	"github.com/gin-gonic/gin"
)

type ReadBookController struct {
	useCase *application.ReadBook
}

func NewReadBookController(uc *application.ReadBook) *ReadBookController {
	return &ReadBookController{useCase: uc}
}

func (c *ReadBookController) Handle(ctx *gin.Context) {
	books, err := c.useCase.ExecuteAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, books)
}
