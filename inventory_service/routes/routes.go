package routes

import (
	"inventory_service/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, productController *controllers.ProductController) {
	api := router.Group("/api/inventory")
	{
		api.GET("/health", func(c *gin.Context) {
			c.String(http.StatusOK, "Inventory Service is healthy")
		})

		api.POST("/products", productController.CreateProduct)
		api.GET("/products", productController.GetAllProducts)
		
		// Concurrency-safe endpoint specific for billing checkout process
		api.POST("/deduct", productController.DeductProducts)
	}
}
