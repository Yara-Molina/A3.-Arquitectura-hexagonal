package controllers

import (
	"net/http"
	"strconv"

	"project1/src/books/application"
	"project1/src/books/domain/entities"

	"github.com/gin-gonic/gin"
)

type UpdateBookController struct {
	useCase *application.UpdateBook
}

func NewUpdateBookController(uc *application.UpdateBook) *UpdateBookController {
	return &UpdateBookController{useCase: uc}
}

func (c *UpdateBookController) Handle(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var book entities.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book.ID = int32(id)

	err = c.useCase.Update(book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Libro actualizado correctamente"})
}
