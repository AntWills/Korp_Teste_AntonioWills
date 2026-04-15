package main

import (
	"log"

	"inventory_service/controllers"
	"inventory_service/infra/database"
	"inventory_service/persistence"
	"inventory_service/routes"
	"inventory_service/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	database.InitDB()

	// Auto migrate product model
	err := database.DB.AutoMigrate(&persistence.Product{})
	if err != nil {
		log.Fatalf("Falha ao rodar as migrações: %v", err)
	}

	// Setup dependency injection
	repo := persistence.NewProductRepository(database.DB)
	service := services.NewProductService(repo)
	controller := controllers.NewProductController(service)

	// Setup Gin router
	router := gin.Default()
	routes.RegisterRoutes(router, controller)

	log.Println("Starting Inventory Service on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
