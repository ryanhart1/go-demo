package database

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToPostgreSQL() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=product-service-db user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	return db, nil
}

func GetDB() *gorm.DB {
	if DB == nil {
		return establishNewConnection()
	}

	psqlDB, err := DB.DB()
	if err != nil {
		return establishNewConnection()
	}

	if err := psqlDB.Ping(); err != nil {
		return establishNewConnection()
	}

	return DB
}

func establishNewConnection() *gorm.DB {
	newDB, err := ConnectToPostgreSQL()
	if err != nil {
		slog.Error("Error establishing new database connection", "error", err.Error())
		return nil
	}
	DB = newDB
	return DB
}

func CloseDB(db *gorm.DB) {
	dbInstance, _ := db.DB()
	err := dbInstance.Close()
	if err != nil {
		slog.Error("Error while closing DB connection.", "error", err.Error())
	} else {
		slog.Info("DB connection closed successfully")
	}
}
