package usecase

import (
	"fmt"
	"go-api/model"
	"go-api/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUSecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (p *ProductUsecase) GetProducts() ([]model.Product, error) {
	return p.repository.GetProducts()

}

func (p *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	productID, err := p.repository.CreateProduct(product)
	if err != nil {
		fmt.Println("Erro ao criar produto no repository:", err)
		return model.Product{}, err
	}
	product.ID = productID
	return product, nil
}
