package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (p *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, name_product, price FROM products"
	rows, err := p.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}
	var productList []model.Product
	var productObject model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObject.ID,
			&productObject.Name,
			&productObject.Price)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}
		productList = append(productList, productObject)
	}
	rows.Close()

	return productList, nil
}
