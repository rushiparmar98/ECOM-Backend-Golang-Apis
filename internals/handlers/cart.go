package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rushi/Desktop/ecom/internals/services"
	"github.com/rushi/Desktop/ecom/models"
)

func AddToCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item models.CartItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := services.AddToCart(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Item added to cart successfully"}`))
}

func GetCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cart := services.GetCart()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cart)
}
