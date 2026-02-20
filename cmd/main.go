package main

import (
	"fmt"
	"go-api/config"
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	config.Load()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//camada repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	//camada usecase
	ProductUsecase := usecase.NewProductUSecase(ProductRepository)

	//camada controller
	ProductController := controller.NewProductController(ProductUsecase)

	server := gin.Default()

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)

	server.Run(fmt.Sprintf(":%d", config.APIPort))

	server.Run(fmt.Sprintf(":%d", config.APIPort))
	fmt.Println("Server running on port", config.APIPort)

}
