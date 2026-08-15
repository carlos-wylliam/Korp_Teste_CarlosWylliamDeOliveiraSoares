package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"faturamento/database"
	"faturamento/models"

	"github.com/gin-gonic/gin"
)

func CriarNota(c *gin.Context) {
	var nota models.Nota
	if err := c.ShouldBindJSON(&nota); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := database.DB.Exec(
		"INSERT INTO notas (status) VALUES (?)",
		"Aberta",
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao salvar nota: " + err.Error()})
		return
	}

	numero, _ := result.LastInsertId()
	nota.Numero = int(numero)
	nota.Status = "Aberta"

	for _, item := range nota.Itens {
		_, err := database.DB.Exec(
			"INSERT INTO itens_nota (nota_numero, produto_codigo, quantidade) VALUES (?, ?, ?)",
			nota.Numero, item.ProdutoCodigo, item.Quantidade,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao salvar item: " + err.Error()})
			return
		}
	}

	c.JSON(http.StatusCreated, nota)
}

func ListarNotas(c *gin.Context) {
	rows, err := database.DB.Query("SELECT numero, status FROM notas")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao buscar notas: " + err.Error()})
		return
	}
	defer rows.Close()

	var notas []models.Nota
	for rows.Next() {
		var n models.Nota
		if err := rows.Scan(&n.Numero, &n.Status); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao ler nota: " + err.Error()})
			return
		}

		itemRows, err := database.DB.Query("SELECT produto_codigo, quantidade FROM itens_nota WHERE nota_numero = ?", n.Numero)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao buscar itens: " + err.Error()})
			return
		}
		for itemRows.Next() {
			var item models.ItemNota
			itemRows.Scan(&item.ProdutoCodigo, &item.Quantidade)
			n.Itens = append(n.Itens, item)
		}
		itemRows.Close()

		notas = append(notas, n)
	}

	c.JSON(http.StatusOK, notas)
}

const estoqueURL = "http://localhost:8000"

func ImprimirNota(c *gin.Context) {
	numero := c.Param("numero")

	var status string
	err := database.DB.QueryRow("SELECT status FROM notas WHERE numero = ?", numero).Scan(&status)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "nota não encontrada"})
		return
	}

	if status != "Aberta" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "somente notas com status Aberta podem ser impressas"})
		return
	}

	rows, err := database.DB.Query("SELECT produto_codigo, quantidade FROM itens_nota WHERE nota_numero = ?", numero)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao buscar itens: " + err.Error()})
		return
	}
	defer rows.Close()

	var itens []models.ItemNota
	for rows.Next() {
		var item models.ItemNota
		rows.Scan(&item.ProdutoCodigo, &item.Quantidade)
		itens = append(itens, item)
	}

	for _, item := range itens {
		body, _ := json.Marshal(map[string]float64{"quantidade": item.Quantidade})
		req, _ := http.NewRequest(http.MethodPut, estoqueURL+"/produtos/"+item.ProdutoCodigo+"/saldo", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "serviço de estoque indisponível: " + err.Error()})
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			c.JSON(http.StatusBadRequest, gin.H{"error": "falha ao atualizar saldo do produto " + item.ProdutoCodigo})
			return
		}
	}

	_, err = database.DB.Exec("UPDATE notas SET status = ? WHERE numero = ?", "Fechada", numero)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao fechar nota: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"numero": numero, "status": "Fechada"})
}