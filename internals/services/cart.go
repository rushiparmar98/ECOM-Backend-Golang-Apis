package services

import (
	"errors"

	"github.com/rushi/Desktop/ecom/internals/daos"
	"github.com/rushi/Desktop/ecom/models"
)

func AddToCart(item models.CartItem) error {
	if item.Quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}
	product, err := daos.GetProductByID(item.ProductID)
	if err != nil {
		return err
	}

	if item.Quantity > product.Quantity {
		return errors.New("requested quantity exceeds available stock")
	}

	daos.AddItemToCart(item)
	return nil
}

// GetCart retrieves the current cart
func GetCart() models.Cart {
	return daos.GetCart()
}
