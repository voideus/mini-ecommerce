package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" binding:"requited"`
	Email    string `json:"email" binding:"required,email" gorm:"unique"`
	Password string `json:"password" binding:"required"`
}
