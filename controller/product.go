package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(productUsecase usecase.ProductUsecase) ProductController {
	return ProductController{
		productUsecase: productUsecase,
	}
}

func (p ProductController) GetProducts(ctx *gin.Context) {
	products, err := p.productUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (p ProductController) CreateProduct(ctx *gin.Context) {
	var product model.Product

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	product, err := p.productUsecase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

func (p ProductController) GetByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	if idParam == "" {
		ctx.JSON(http.StatusBadRequest, "ID do produto não pode ser nulo")
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "ID do produto precisa ser um número")
		return
	}

	product, err := p.productUsecase.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p ProductController) DeleteProduct(ctx *gin.Context) {
	idParam := ctx.Param("id")
	if idParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID do produto não pode ser vazio"})
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID precisa ser um número"})
		return
	}

	err = p.productUsecase.DeleteProduct(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil) // REST mais correto para DELETE
}
