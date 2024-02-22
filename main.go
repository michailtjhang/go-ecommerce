package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Middleware: Logger
	router.Use(gin.Logger())

	// Middleware: Recovery
	router.Use(gin.Recovery())

	// Routes definition
	router.GET("/hello", func(c *gin.Context) {
		// 2xx => success
		// 3xx => redirect
		// 4xx => bad request /client error
		// 5xx => internal server error
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// Route parameters / dinamic routes
	router.GET("/halo/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "Halo, " + name + "!",
		})
	})

	router.POST("/login", func(c *gin.Context) {
		var loginData struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		// Di sini Anda dapat melakukan validasi login, misalnya memeriksa di database, dll
		// contoh sederhana: Memeriksa apakah username dan password cocok
		if loginData.Username == "admin" && loginData.Password == "password" {
			c.JSON(200, gin.H{
				"message": "Login successful",
			})
		} else {
			c.JSON(401, gin.H{
				"message": "Invalid credentials",
			})
		}
	})

	// menambahkan endpoint untuk mengembalikan parameter query
	router.GET("/users", func(c *gin.Context) {
		name := c.Query("name")

		if name == "" {
			c.JSON(400, gin.H{
				"error": "Name parameter is missing",
			})
			return
		} else {
			c.JSON(200, gin.H{
				"message": "Hello, " + name + "!",
			})
		}
	})

	// jalankan server
	router.Run(":8080")

}
