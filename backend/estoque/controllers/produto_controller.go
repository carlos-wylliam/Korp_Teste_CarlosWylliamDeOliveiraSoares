package controllers

import (
	"net/http"

	"estoque/database"
	"estoque/models"

	"github.com/gin-gonic/gin"
)

func CriarProduto(c *gin.Context) {
	var produto models.Produto
	if err := c.ShouldBindJSON(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := database.DB.Exec(
		"INSERT INTO produtos (codigo, descricao, saldo) VALUES (?, ?, ?)",
		produto.Codigo, produto.Descricao, produto.Saldo,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao salvar produto: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, produto)
}

func ListarProdutos(c *gin.Context) {
	rows, err := database.DB.Query("SELECT codigo, descricao, saldo FROM produtos")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao buscar produtos: " + err.Error()})
		return
	}
	defer rows.Close()

	var produtos []models.Produto
	for rows.Next() {
		var p models.Produto
		if err := rows.Scan(&p.Codigo, &p.Descricao, &p.Saldo); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao ler produto: " + err.Error()})
			return
		}
		produtos = append(produtos, p)
	}

	c.JSON(http.StatusOK, produtos)
}