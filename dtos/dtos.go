package dtos

import "github.com/rushi/Desktop/ecom/models"

type CartResponse struct {
	Items []models.CartItem `json:"items"`
}
