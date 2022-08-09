package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	User      User `gorm:"foreignkey:UserID"`
	UserID    uint
	Product   Product `gorm:"foreignkey:ProductID"`
	ProductID uint
	Quantity  int `json:"quantity"`
}
