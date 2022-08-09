package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Quantity    int    `json:"quantity"`
	Description string `json:"description"`
}
