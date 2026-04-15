package controllers

import (
	"billing_service/services"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InvoiceController struct {
	service services.InvoiceService
}

func NewInvoiceController(service services.InvoiceService) *InvoiceController {
	return &InvoiceController{service: service}
}

// CreateInvoice godoc
// POST /api/billing/invoices
func (c *InvoiceController) CreateInvoice(ctx *gin.Context) {
	var input services.CreateInvoiceInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "payload inválido: " + err.Error()})
		return
	}

	invoice, err := c.service.CreateInvoice(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, invoice)
}

// GetAllInvoices godoc
// GET /api/billing/invoices
func (c *InvoiceController) GetAllInvoices(ctx *gin.Context) {
	invoices, err := c.service.GetAllInvoices()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao buscar notas fiscais"})
		return
	}

	ctx.JSON(http.StatusOK, invoices)
}

// PrintInvoice godoc
// POST /api/billing/invoices/:id/print
//
// Mapeamento de erros de domínio → HTTP:
//   - ErrInvoiceNotFound      → 404 Not Found
//   - ErrInvoiceNotOpen       → 409 Conflict  (estado inválido, semântica mais precisa que 400)
//   - ErrInventoryUnavailable → 503 Service Unavailable
//   - outros                  → 500 Internal Server Error
func (c *InvoiceController) PrintInvoice(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido: deve ser um número inteiro positivo"})
		return
	}

	invoice, err := c.service.PrintInvoice(uint(id))
	if err != nil {
		switch {
		case errors.Is(err, services.ErrInvoiceNotFound):
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case errors.Is(err, services.ErrInvoiceNotOpen):
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		case errors.Is(err, services.ErrInventoryUnavailable):
			ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "erro interno: " + err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, invoice)
}
