package main

import (
	"github.com/Didinoer/go-rest-api.git/controllers"
	"github.com/Didinoer/go-rest-api.git/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi router Gin
	r := gin.Default()

	// Koneksi ke database
	models.ConnectDatabase()

	// Buat instance dari controller
	productController := controllers.ProductsController{}

	// Tambahkan endpoint
	r.GET("api/products", productController.Index)
	r.GET("api/products/:id", productController.Show)
	r.GET("api/products/nama/:id", productController.Shownama)

	// Jalankan server
	r.Run()
}
