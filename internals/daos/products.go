package daos

import (
	"github.com/rushi/Desktop/ecom/configs"
	"github.com/rushi/Desktop/ecom/models"
)

// AddProduct adds a product to the in-memory storage
func AddProduct(product models.Product) (models.Product, error) {
	query := `INSERT INTO products (name, price, quantity) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := configs.DB.QueryRow(query, product.Name, product.Price, product.Quantity).Scan(&id)
	if err != nil {
		return models.Product{}, err
	}
	product.ID = id
	return product, nil
}

// ListProducts retrieves all products from storage
func ListProducts() ([]models.Product, error) {
	rows, err := configs.DB.Query("SELECT id, name, price, quantity FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

// GetProductByID retrieves a product by its ID
func GetProductByID(productID int) (models.Product, error) {
	var product models.Product
	query := "SELECT id, name, price, quantity FROM products WHERE id = $1"
	err := configs.DB.QueryRow(query, productID).Scan(&product.ID, &product.Name, &product.Price, &product.Quantity)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

// UpdateProduct updates a product's quantity in storage
func UpdateProduct(product models.Product) error {
	query := "UPDATE products SET name = $1, price = $2, quantity = $3 WHERE id = $4"
	_, err := configs.DB.Exec(query, product.Name, product.Price, product.Quantity, product.ID)
	return err
}
