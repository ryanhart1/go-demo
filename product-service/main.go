package main

import (
	"log"
	"net/http"
	"os"
	"product-service/database"
	"product-service/models"

	"gorm.io/gorm/logger"
)

func main() {
	db, _ := database.ConnectToPostgreSQL()

	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running migrations")

	// Migrate the schema
	db.AutoMigrate(&models.Product{})

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/products", getProducts)
	http.HandleFunc("/products/{id}", getProduct)
	http.HandleFunc("/createproduct", createProduct)
	http.HandleFunc("/deleteproduct/{id}", deleteProduct)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("Failed to start server. \n", err)
		os.Exit(2)
	}

}
