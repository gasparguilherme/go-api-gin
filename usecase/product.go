package usecase

import (
	"go-api/model"
)

type ProductUsecase struct {
	//repository
}

func NewProductUSecase() ProductUsecase {
	return ProductUsecase{}
}

func (p *ProductUsecase) GetProducts() ([]model.Product, error) {
	return []model.Product{}, nil
}
