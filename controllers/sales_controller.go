package controllers

import (
	"encoding/json"
	"inventory-management-system/config"
	"inventory-management-system/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func CreateSalesOrder(w http.ResponseWriter, r *http.Request) {
	var order models.SalesOrder
	json.NewDecoder(r.Body).Decode(&order)
	order.Date = time.Now()
	config.DB.Create(&order)

	// Update inventory
	var product models.Product
	config.DB.First(&product, order.ProductID)
	product.Quantity -= order.Quantity
	config.DB.Save(&product)

	json.NewEncoder(w).Encode(order)
}

func GetSalesOrders(w http.ResponseWriter, r *http.Request) {
	var orders []models.SalesOrder
	config.DB.Find(&orders)
	json.NewEncoder(w).Encode(orders)
}

func GetSalesOrder(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var order models.SalesOrder
	config.DB.First(&order, id)
	json.NewEncoder(w).Encode(order)
}
