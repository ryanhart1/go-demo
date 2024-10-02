package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"product-service/database"
	"product-service/models"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is the product service!\n")
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	database.GetDB().Find(&products)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	id := r.PathValue("id")
	result := database.GetDB().First(&product, id)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Could not find any products with that ID")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("%+v\n%#v\n", product, product)
	database.GetDB().Create(&product)
	json.NewEncoder(w).Encode(product)

}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Printf("The id is : %s", id)
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Please provide an id")
		return
	}

	product := &models.Product{}
	result := database.GetDB().Delete(product, id)
	if result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Could not delete product")
		return
	} else if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Could not find any products with that ID")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Product deleted")
}
