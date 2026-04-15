package controllers

import (
	"inventory_service/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service *services.ProductService
}

func NewProductController(service *services.ProductService) *ProductController {
	return &ProductController{service: service}
}

type CreateProductRequest struct {
	Code        string `json:"code" binding:"required"`
	Description string `json:"description" binding:"required"`
	Balance     int    `json:"balance" binding:"required,gte=0"`
}

func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var req CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := c.service.CreateProduct(req.Code, req.Description, req.Balance)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

func (c *ProductController) GetAllProducts(ctx *gin.Context) {
	products, err := c.service.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

type DeductRequest struct {
	Items []services.DeductItem `json:"items" binding:"required,dive"`
}

func (c *ProductController) DeductProducts(ctx *gin.Context) {
	var req DeductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.DeductProducts(req.Items)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Estoque deduzido com sucesso"})
}
