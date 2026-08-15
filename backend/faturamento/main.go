package main

import (
	"faturamento/controllers"
	"faturamento/database"

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

	r.POST("/notas", controllers.CriarNota)
	r.GET("/notas", controllers.ListarNotas)

	r.Run(":8001")
}