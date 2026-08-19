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

	var existe int
	database.DB.QueryRow("SELECT COUNT(*) FROM produtos WHERE codigo = ?", produto.Codigo).Scan(&existe)
	if existe > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "já existe um produto com esse código"})
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

type AtualizarSaldoRequest struct {
	Quantidade float64 `json:"quantidade"`
}

func AtualizarSaldo(c *gin.Context) {
	codigo := c.Param("codigo")

	var req AtualizarSaldoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var saldoAtual float64
	err := database.DB.QueryRow("SELECT saldo FROM produtos WHERE codigo = ?", codigo).Scan(&saldoAtual)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "produto não encontrado"})
		return
	}

	if saldoAtual < req.Quantidade {
		c.JSON(http.StatusBadRequest, gin.H{"error": "saldo insuficiente para o produto " + codigo})
		return
	}

	novoSaldo := saldoAtual - req.Quantidade
	_, err = database.DB.Exec("UPDATE produtos SET saldo = ? WHERE codigo = ?", novoSaldo, codigo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao atualizar saldo: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"codigo": codigo, "novoSaldo": novoSaldo})
}

func BuscarProduto(c *gin.Context) {
	codigo := c.Param("codigo")

	var produto models.Produto
	err := database.DB.QueryRow("SELECT codigo, descricao, saldo FROM produtos WHERE codigo = ?", codigo).
		Scan(&produto.Codigo, &produto.Descricao, &produto.Saldo)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "produto não encontrado"})
		return
	}

	c.JSON(http.StatusOK, produto)
}

func ExcluirProduto(c *gin.Context) {
	codigo := c.Param("codigo")

	result, err := database.DB.Exec("DELETE FROM produtos WHERE codigo = ?", codigo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao excluir produto: " + err.Error()})
		return
	}

	linhasAfetadas, _ := result.RowsAffected()
	if linhasAfetadas == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "produto não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "produto excluído com sucesso"})
}