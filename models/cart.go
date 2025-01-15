package models

// CartItem represents a product in the shopping cart
type CartItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

// Cart holds the cart items for the user
type Cart struct {
	UserID int        `json:"user_id"`
	Items  []CartItem `json:"items"`
}
