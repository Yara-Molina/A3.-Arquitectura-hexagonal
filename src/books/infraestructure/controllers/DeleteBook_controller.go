package controllers

import (
	"net/http"
	"strconv"

	"project1/src/books/application"

	"github.com/gin-gonic/gin"
)

type DeleteBookController struct {
	useCase *application.DeleteBook
}

func NewDeleteBookController(useCase *application.DeleteBook) *DeleteBookController {
	return &DeleteBookController{useCase: useCase}
}

func (c *DeleteBookController) Execute(ctx *gin.Context) {
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

	ctx.JSON(http.StatusOK, gin.H{"message": "Libro eliminado correctamente"})
}
