package handlers

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root@tcp(localhost:3306)/ecomm?charset=utf8mb4&parseTime=True&loc=Local") 
	if err != nil {
		panic("failed to connect database")
	}

	// db.AutoMigrate(&models.TransactionItem{}) // Migrate model yang diperlukan
}