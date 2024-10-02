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
	// dsn := fmt.Sprintf(
	// 	"host=localhost user=%s password=%s dbname=%s port=5433 sslmode=disable",
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_NAME_ORDER"),
	// )
	dsn := fmt.Sprintf(
		"host=orderdb user=ryan password=admin dbname=adminorder port=5432 sslmode=disable",
	)
	// dsn := fmt.Sprintf(
	// 	"host=localhost user=ryan password=admin dbname=admin port=5432 sslmode=disable",
	// )

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	return db, nil
}

func GetDB() *gorm.DB {
	// Check if DB is initialized
	if DB == nil {
		// Try to establish a new connection if DB is nil
		return establishNewConnection()
	}

	// Get the underlying sql.DB object from the GORM DB instance
	psqlDB, err := DB.DB()
	if err != nil {
		// If error in getting the sql.DB, try to establish a new connection
		return establishNewConnection()
	}

	// Ping the database to check the connection
	if err := psqlDB.Ping(); err != nil {
		// If ping fails, try to establish a new connection
		return establishNewConnection()
	}

	// Return the existing DB instance if all checks pass
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
		slog.Error("Error while closing DB connection. Not a problem actually", "error", err.Error())
	} else {
		slog.Info("DB connection is closed successfully")
	}
}
