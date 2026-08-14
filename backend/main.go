package main

import (
	"korpteste-backend/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/produtos", controllers.CriarProduto)

	r.Run(":8000")
}