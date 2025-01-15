package models

import "time"

type Order struct {
	ID          int        `json:"id"`
	UserID      int        `json:"user_id"`
	TotalAmount float64    `json:"total_amount"` // This is important
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Items       []CartItem `json:"items"`
}
