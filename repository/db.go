package repository

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/voideus/mini-ecommerce/model"
)

func DB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("mini-ecommerce"), &gorm.Config{})
	if err != nil {
		fmt.Println("Database connection error: ", err.Error())
		panic("Database connection failed")
	}

	db.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{})
	return db
}
