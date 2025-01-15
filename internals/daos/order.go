package daos

import (
	"database/sql"

	"github.com/rushi/Desktop/ecom/configs"
	"github.com/rushi/Desktop/ecom/models"
)

// AddOrder adds a new order to the storage
func AddOrder(order models.Order) (models.Order, error) {
	query := `INSERT INTO orders (user_id, status, total_amount) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := configs.DB.QueryRow(query, order.UserID, order.Status, order.TotalAmount).Scan(&id)
	if err != nil {
		return models.Order{}, err
	}
	order.ID = id

	// Insert items into the order_items table
	for _, item := range order.Items {
		_, err := configs.DB.Exec("INSERT INTO order_items (order_id, product_id, quantity) VALUES ($1, $2, $3)", order.ID, item.ProductID, item.Quantity)
		if err != nil {
			return models.Order{}, err
		}
	}

	return order, nil
}

// GetOrderByID retrieves an order by its ID
func GetOrderByID(id int) (models.Order, error) {

	var order models.Order
	query := `SELECT id, user_id, status, total_amount FROM orders WHERE id = $1`
	row := configs.DB.QueryRow(query, id)

	err := row.Scan(&order.ID, &order.UserID, &order.Status, &order.TotalAmount)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Order{}, sql.ErrNoRows
		}
		return models.Order{}, err
	}

	// Fetch items for the order
	itemRows, err := configs.DB.Query("SELECT product_id, quantity FROM order_items WHERE order_id = $1", order.ID)
	if err != nil {
		return models.Order{}, err
	}
	defer itemRows.Close()

	var items []models.CartItem
	for itemRows.Next() {
		var item models.CartItem
		if err := itemRows.Scan(&item.ProductID, &item.Quantity); err != nil {
			return models.Order{}, err
		}
		items = append(items, item)
	}
	order.Items = items

	return order, nil
}

// UpdateOrder updates an existing order
func UpdateOrder(orderID int, status string) error {
	query := "UPDATE orders SET status = $1 WHERE id = $2"
	_, err := configs.DB.Exec(query, status, orderID)
	return err
}

// ListOrders retrieves all orders from storage
func ListOrders() ([]models.Order, error) {
	rows, err := configs.DB.Query("SELECT id, user_id, status, total_amount FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.ID, &order.UserID, &order.Status, &order.TotalAmount); err != nil {
			return nil, err
		}

		// Get the items for each order
		itemRows, err := configs.DB.Query("SELECT product_id, quantity FROM order_items WHERE order_id = $1", order.ID)
		if err != nil {
			return nil, err
		}
		defer itemRows.Close()

		var items []models.CartItem
		for itemRows.Next() {
			var item models.CartItem
			if err := itemRows.Scan(&item.ProductID, &item.Quantity); err != nil {
				return nil, err
			}
			items = append(items, item)
		}
		order.Items = items

		orders = append(orders, order)
	}

	return orders, nil
}
