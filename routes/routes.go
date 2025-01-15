package routes

import (
	"github.com/gorilla/mux"
	"github.com/rushi/Desktop/ecom/internals/handlers"
)

func SetupRoutes(r *mux.Router) {

	// Product Routes
	r.HandleFunc("/product", handlers.AddProduct).Methods("POST")
	r.HandleFunc("/products", handlers.ListProducts).Methods("GET")

	// Cart Routes
	r.HandleFunc("/cart", handlers.AddToCart).Methods("POST") // Add to cart
	r.HandleFunc("/cart", handlers.GetCart).Methods("GET")    // Get cart

	// Order Routes
	r.HandleFunc("/order", handlers.PlaceOrder).Methods("POST")
	r.HandleFunc("/order/{id}/cancel", handlers.CancelOrder).Methods("POST")
	r.HandleFunc("/order/{id}/update", handlers.UpdateOrder).Methods("PUT")
}
