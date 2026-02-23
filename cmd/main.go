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

	ProductRepository := repository.NewProductRepository(dbConnection)

	ProductUsecase := usecase.NewProductUSecase(ProductRepository)

	ProductController := controller.NewProductController(ProductUsecase)

	server := gin.Default()

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:id", ProductController.GetByID)

	server.Run(fmt.Sprintf(":%d", config.APIPort))

	fmt.Println("Server running on port", config.APIPort)

}
