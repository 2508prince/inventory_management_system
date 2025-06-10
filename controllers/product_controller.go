package controllers

import (
	"encoding/json"
	"inventory-management-system/config"
	"inventory-management-system/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)
	config.DB.Create(&product)
	json.NewEncoder(w).Encode(product)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	config.DB.Find(&products)
	json.NewEncoder(w).Encode(products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var product models.Product
	result := config.DB.First(&product, id)
	if result.Error != nil {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var product models.Product
	config.DB.First(&product, id)
	json.NewDecoder(r.Body).Decode(&product)
	config.DB.Save(&product)
	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	config.DB.Delete(&models.Product{}, id)
	w.WriteHeader(http.StatusNoContent)
}
