package main

import (
	"fmt"
	"inventory-management-system/config"
	"inventory-management-system/models"
	"inventory-management-system/routes"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	config.DB.AutoMigrate(
		&models.Product{},
		&models.SalesOrder{},
		&models.PurchaseOrder{},
	)

	router := routes.RegisterRoutes()

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
