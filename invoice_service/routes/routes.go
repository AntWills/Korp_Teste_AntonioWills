package routes

import (
	"billing_service/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRoutes registra todas as rotas do billing_service no grupo /api/billing.
func SetupRoutes(router *gin.Engine, invoiceCtrl *controllers.InvoiceController) {
	api := router.Group("/api/invoices")
	{
		api.GET("/health", func(c *gin.Context) {
			c.String(http.StatusOK, "Invoices Service is healthy")
		})

		api.POST("", invoiceCtrl.CreateInvoice)
		api.GET("", invoiceCtrl.GetAllInvoices)
		api.POST("/:id/print", invoiceCtrl.PrintInvoice)

	}
}
