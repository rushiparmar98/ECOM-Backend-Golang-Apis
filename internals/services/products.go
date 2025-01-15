package services

import (
	"github.com/rushi/Desktop/ecom/internals/daos"
	"github.com/rushi/Desktop/ecom/models"
)

// AddProduct adds a new product
func AddProduct(product models.Product) (models.Product, error) {
	return daos.AddProduct(product)
}

// ListProducts retrieves all products
func ListProducts() ([]models.Product, error) {
	return daos.ListProducts()
}
