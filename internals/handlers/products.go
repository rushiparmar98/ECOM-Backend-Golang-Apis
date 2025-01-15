package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rushi/Desktop/ecom/internals/services"
	"github.com/rushi/Desktop/ecom/models"
)

// AddProductHandler adds a new product
func AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if product.Name == "" || product.Price <= 0 || product.Quantity <= 0 {
		http.Error(w, "Invalid product details", http.StatusBadRequest)
		return
	}

	createdProduct, err := services.AddProduct(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdProduct)
}

// ListProductsHandler lists all products
func ListProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products, err := services.ListProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
