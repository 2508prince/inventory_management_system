package routes

import (
	"inventory-management-system/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()

	// Product routes
	router.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", controllers.GetProduct).Methods("GET")
	router.HandleFunc("/products/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", controllers.DeleteProduct).Methods("DELETE")

	// Sales order routes
	router.HandleFunc("/sales", controllers.CreateSalesOrder).Methods("POST")
	router.HandleFunc("/sales", controllers.GetSalesOrders).Methods("GET")
	router.HandleFunc("/sales/{id}", controllers.GetSalesOrder).Methods("GET")

	// Purchase order routes
	router.HandleFunc("/purchase", controllers.CreatePurchaseOrder).Methods("POST")
	router.HandleFunc("/purchase", controllers.GetPurchaseOrders).Methods("GET")
	router.HandleFunc("/purchase/{id}", controllers.GetPurchaseOrder).Methods("GET")

	// Inventory route
	router.HandleFunc("/inventory", controllers.GetInventory).Methods("GET")

	return router
}
