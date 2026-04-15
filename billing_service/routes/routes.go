package routes

import (
	"billing_service/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRoutes registra todas as rotas do billing_service no grupo /api/billing.
func SetupRoutes(router *gin.Engine, invoiceCtrl *controllers.InvoiceController) {
	api := router.Group("/api/billing")
	{
		api.GET("/health", func(c *gin.Context) {
			c.String(http.StatusOK, "Billing Service is healthy")
		})

		invoices := api.Group("/invoices")
		{
			invoices.POST("", invoiceCtrl.CreateInvoice)
			invoices.GET("", invoiceCtrl.GetAllInvoices)
			invoices.POST("/:id/print", invoiceCtrl.PrintInvoice)
		}
	}
}
