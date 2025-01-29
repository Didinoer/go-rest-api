// package controllers

// import (
// 	"net/http"

// 	"github.com/Didinoer/go-rest-api.git/models"
// 	"github.com/gin-gonic/gin"
// )

// func index(c *gin.Context) {
// 	var product []models.Product

// 	models.DB.Find(&product)

// 	c.JSON(http.StatusOK, gin.H{"products": product})
// }

package controllers

import (
	"net/http"

	"github.com/Didinoer/go-rest-api.git/models"
	"github.com/gin-gonic/gin"
)

// Struct untuk ProductsController
type ProductsController struct{}

// Method Index untuk menangani GET /api/products
func (pc ProductsController) Index(c *gin.Context) {
	var products []models.Product

	// Ambil semua data produk dari database
	models.DB.Find(&products)

	// Kirimkan data produk dalam format JSON
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func (pc ProductsController) Show(c *gin.Context) {
	var products []models.Product

	// Ambil semua data produk dari database
	models.DB.Find(&products)

	// Kirimkan data produk dalam format JSON
	c.JSON(http.StatusOK, gin.H{"products": products})
}
