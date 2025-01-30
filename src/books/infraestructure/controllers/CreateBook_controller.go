package controllers

import (
	"net/http"
	"project1/src/books/application"
	"project1/src/books/domain/entities"

	"github.com/gin-gonic/gin"
)

type CreateBookController struct {
	useCase *application.CreateBook
}

func NewCreateBookController(uc *application.CreateBook) *CreateBookController {
	return &CreateBookController{useCase: uc}
}

func (c *CreateBookController) Handle(ctx *gin.Context) {
	var newBook entities.Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	book, err := c.useCase.Execute(newBook)
	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, book)
}
