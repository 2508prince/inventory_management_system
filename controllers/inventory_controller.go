package controllers

import (
	"encoding/json"
	"inventory-management-system/config"
	"inventory-management-system/models"
	"net/http"
)

func GetInventory(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	config.DB.Find(&products)

	var inventory []models.Inventory
	for _, p := range products {
		inventory = append(inventory, models.Inventory{
			ProductID: p.ID,
			Quantity:  p.Quantity,
		})
	}
	json.NewEncoder(w).Encode(inventory)
}
