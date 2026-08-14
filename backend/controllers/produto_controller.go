package controllers

import (
	"net/http"

	"korpteste-backend/models"

	"github.com/gin-gonic/gin"
)

var produtos []models.Produto

func CriarProduto(c *gin.Context) {
	var produto models.Produto
	if err := c.ShouldBindJSON(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	produtos = append(produtos, produto)
	c.JSON(http.StatusCreated, produto)
}

func ListarProdutos(c *gin.Context) {
	c.JSON(http.StatusOK, produtos)
}