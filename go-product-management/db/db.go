package db

import (
	"github.com/akashyap17/go-product-management/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=admin password=admin dbname=product_management port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate the schema
	err = DB.AutoMigrate(&models.Category{}, &models.Product{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
