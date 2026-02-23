package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"go-api/config"
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"
)

func main() {
	// Carrega configurações
	config.Load()

	// Conecta ao banco de dados
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Inicializa repository, usecase e controller
	ProductRepository := repository.NewProductRepository(dbConnection)
	ProductUsecase := usecase.NewProductUSecase(ProductRepository)
	ProductController := controller.NewProductController(ProductUsecase)

	// Cria o servidor Gin
	server := gin.Default()

	// Middleware CORS (permite todas as origens — ótimo para teste local)
	server.Use(cors.Default())

	// Rotas
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:id", ProductController.GetByID)
	server.DELETE("/product/:id", ProductController.DeleteProduct)
	// Mensagem antes de rodar o servidor
	fmt.Println("Server running on port", config.APIPort)

	// Roda o servidor
	server.Run(fmt.Sprintf(":%d", config.APIPort))
}
