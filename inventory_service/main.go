package main

import (
	"fmt"
	"log"
	"net/http"

	"inventory_service/controllers"
	"inventory_service/infra/database"
	"inventory_service/persistence"
	"inventory_service/routes"
	"inventory_service/services"
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

	// Setup routes
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux, controller)

	fmt.Println("Starting Inventory Service on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
