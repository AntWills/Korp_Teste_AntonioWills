package main

import (
	"billing_service/controllers"
	"billing_service/infra/database"
	"billing_service/persistence"
	"billing_service/routes"
	"billing_service/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	err := database.DB.AutoMigrate(&persistence.Invoice{}, &persistence.InvoiceItem{})
	if err != nil {
		log.Fatalf("Falha ao rodar as migrações: %v", err)
	}

	repo := persistence.NewInvoiceRepository(database.DB)
	inventoryClient := services.NewInventoryClient()

	service := services.NewInvoiceService(repo, inventoryClient)

	// Controller
	invoiceCtrl := controllers.NewInvoiceController(service)

	router := gin.Default()
	routes.SetupRoutes(router, invoiceCtrl)

	log.Println("Starting Billing Service on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
