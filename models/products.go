package models

type Product struct {
	ID       int     `json:"id" gorm:"primary_key"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity" gorm:"schema:your_schema"`
}
