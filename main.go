package main

import (
	"go-ecommerce/configs"
	"go-ecommerce/handlers"
	"go-ecommerce/middlewares"
	"go-ecommerce/migrations"

	"github.com/gin-gonic/gin"

	"net/http"
	_ "net/http/pprof"
)

func main() {
	db, err := configs.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	migrations.Migrate(db)
	// seeders.Seed(db)

	router := gin.Default()

	// Rute CRUD Produk
	router.GET("/products", middlewares.AuthMiddleware(), handlers.ListProducts(db))
	router.GET("/products/:id", handlers.GetProduct(db))
	router.POST("/products", handlers.CreateProduct(db))
	router.PUT("/products/:id", handlers.UpdateProduct(db))
	router.DELETE("/products/:id", handlers.DeleteProduct(db))

	router.POST("/login", handlers.Login(db))
	router.POST("/register", handlers.Register(db))

	router.GET("/debug/pprof/*pprof", gin.WrapH(http.DefaultServeMux))

	// jalankan server
	router.Run(":5000")

}
