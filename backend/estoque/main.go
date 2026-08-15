package main

import (
	"estoque/controllers"
	"estoque/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Conectar()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:4200"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))

	r.POST("/produtos", controllers.CriarProduto)
	r.GET("/produtos", controllers.ListarProdutos)
	r.PUT("/produtos/:codigo/saldo", controllers.AtualizarSaldo)
	
	r.Run(":8000")
}