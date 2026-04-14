package routes

import (
	"inventory_service/controllers"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, productController *controllers.ProductController) {
	// Health Check
	mux.HandleFunc("/api/inventory/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Inventory Service is healthy"))
	})

	// Products
	mux.HandleFunc("/api/inventory/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			productController.GetAllProducts(w, r)
		case http.MethodPost:
			productController.CreateProduct(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
