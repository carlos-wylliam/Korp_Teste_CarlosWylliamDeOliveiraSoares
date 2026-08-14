package main

import (
	"korpteste-backend/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:4200"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))

	r.POST("/produtos", controllers.CriarProduto)
	r.GET("/produtos", controllers.ListarProdutos)
	r.Run(":8000")
}