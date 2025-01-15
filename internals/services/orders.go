package services

import (
	"database/sql"
	"errors"

	"github.com/rushi/Desktop/ecom/internals/daos"
	"github.com/rushi/Desktop/ecom/models"
)

// PlaceOrder processes the current cart into a new order
func PlaceOrder() (models.Order, error) {
	cart := daos.GetCart()
	if len(cart.Items) == 0 {
		return models.Order{}, errors.New("cart is empty")
	}

	totalAmount := 0.0
	for _, item := range cart.Items {
		product, err := daos.GetProductByID(item.ProductID)
		if err != nil {
			return models.Order{}, err
		}

		if item.Quantity > product.Quantity {
			return models.Order{}, errors.New("insufficient stock for product: " + product.Name)
		}

		product.Quantity -= item.Quantity
		if err := daos.UpdateProduct(product); err != nil {
			return models.Order{}, err
		}

		totalAmount += float64(item.Quantity) * product.Price
	}

	order := models.Order{
		UserID:      cart.UserID,
		Status:      "Placed",
		TotalAmount: totalAmount,
	}

	order, err := daos.AddOrder(order)
	if err != nil {
		return models.Order{}, err
	}
	daos.ClearCart()
	return order, nil
}

// CancelOrder changes the status of an existing order to "Cancelled"
func CancelOrder(orderID int) (models.Order, error) {
	order, err := daos.GetOrderByID(orderID)
	if err != nil {
		return models.Order{}, err
	}

	order.Status = "Cancelled"

	if err := daos.UpdateOrder(order.ID, "Cancelled"); err != nil {
		return models.Order{}, err
	}

	return order, nil
}

// UpdateOrder updates the status or items of an existing order
func UpdateOrder(orderID int, updateOrder models.Order) (models.Order, error) {
	order, err := daos.GetOrderByID(orderID)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Order{}, errors.New("order not found")
		}
		return models.Order{}, err
	}

	if updateOrder.Status != "" {
		order.Status = updateOrder.Status
	}

	if len(updateOrder.Items) > 0 {
		order.Items = updateOrder.Items
	}

	if err := daos.UpdateOrder(order.ID, order.Status); err != nil {
		return models.Order{}, err
	}

	return order, nil
}

// ListOrders returns all orders in the system
func ListOrders() ([]models.Order, error) {
	orders, err := daos.ListOrders()
	if err != nil {
		return nil, err
	}
	return orders, nil
}
