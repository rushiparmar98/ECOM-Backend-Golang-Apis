package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rushi/Desktop/ecom/configs"
	"github.com/rushi/Desktop/ecom/internals/services"
	"github.com/rushi/Desktop/ecom/models"
)

func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	order, err := services.PlaceOrder()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func CancelOrder(w http.ResponseWriter, r *http.Request) {
	orderID := mux.Vars(r)["id"]
	log.Printf("Canceling order with ID: %s", orderID)

	// Check if the order exists
	var order models.Order
	err := configs.DB.QueryRow("SELECT id, user_id, status, total_amount, created_at, updated_at FROM orders WHERE id = $1", orderID).Scan(&order.ID, &order.UserID, &order.Status, &order.TotalAmount, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Order not found", http.StatusNotFound)
			log.Printf("Order with ID %s not found", orderID)
			return
		}
		http.Error(w, "Error retrieving order", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Proceed with canceling the order if it exists
	_, err = configs.DB.Exec("UPDATE orders SET status = $1 WHERE id = $2", "Cancelled", orderID)
	if err != nil {
		http.Error(w, "Error canceling order", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	var updateOrder models.Order
	if err := json.NewDecoder(r.Body).Decode(&updateOrder); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	order, err := services.UpdateOrder(id, updateOrder)
	if err != nil {
		fmt.Println("Error updating order:", err)
		if err == sql.ErrNoRows {
			http.Error(w, "Order not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}
