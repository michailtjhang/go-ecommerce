package main

import (
	"go-ecommerce/configs"
	"go-ecommerce/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := configs.InitDB()
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

	router.POST("/login", handlers.Login(db))
	router.POST("/register", handlers.Register(db))

	// jalankan server
	router.Run(":5000")

}
