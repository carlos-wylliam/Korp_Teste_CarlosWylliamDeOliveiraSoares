package controllers

import (
	"net/http"

	"korpteste-backend/models"

	"github.com/gin-gonic/gin"
)

var notas []models.Nota
var proximoNumero = 1

func CriarNota(c *gin.Context) {
	var nota models.Nota
	if err := c.ShouldBindJSON(&nota); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	nota.Numero = proximoNumero
	nota.Status = "Aberta"
	proximoNumero++

	notas = append(notas, nota)
	c.JSON(http.StatusCreated, nota)
}

func ListarNotas(c *gin.Context) {
	c.JSON(http.StatusOK, notas)
}