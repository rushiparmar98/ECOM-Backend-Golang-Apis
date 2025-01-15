package daos

import (
	"fmt"

	"github.com/rushi/Desktop/ecom/configs"
	"github.com/rushi/Desktop/ecom/models"
)

// GetCart retrieves the current cart
func GetCart() models.Cart {
	var cart models.Cart
	rows, err := configs.DB.Query("SELECT product_id, quantity FROM cart")
	if err != nil {
		fmt.Println("Error querying cart:", err)
		return cart
	}
	defer rows.Close()

	for rows.Next() {
		var item models.CartItem
		if err := rows.Scan(&item.ProductID, &item.Quantity); err != nil {
			fmt.Println("Error scanning cart item:", err)
			continue
		}
		cart.Items = append(cart.Items, item)
	}

	if len(cart.Items) == 0 {
		fmt.Println("Cart is empty")
	}

	return cart
}

// AddItemToCart adds an item to the cart
func AddItemToCart(item models.CartItem) {
	_, err := configs.DB.Exec("INSERT INTO cart (product_id, quantity) VALUES ($1, $2)", item.ProductID, item.Quantity)
	if err != nil {
		fmt.Println("Error adding item to cart:", err)
	}
}

// ClearCart removes all items from the cart
func ClearCart() {
	_, err := configs.DB.Exec("DELETE FROM cart")
	if err != nil {
		fmt.Println("Error clearing the cart:", err)
	}
}
