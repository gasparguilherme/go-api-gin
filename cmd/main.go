package main

import (
	"fmt"
	"go-api/config"
	"go-api/controller"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()
	server := gin.Default()

	productUsecase := usecase.NewProductUSecase()
	productController := controller.NewProductController(productUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", productController.GetProducts)

	server.Run(fmt.Sprintf(":%d", config.APIPort))
}
