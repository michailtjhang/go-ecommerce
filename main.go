package main

import (
	"go-ecommerce/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()	

	router := gin.Default()

	// Rute CRUD Produk
	router.GET("/products", handlers.ListProducts(db))
	router.GET("/products/:id", handlers.GetProduct(db))
	router.POST("/products", handlers.CreateProduct(db))
	router.PUT("/products/:id", handlers.UpdateProduct(db))
	router.DELETE("/products/:id", handlers.DeleteProduct(db))

	// jalankan server
	router.Run(":8080")

}
