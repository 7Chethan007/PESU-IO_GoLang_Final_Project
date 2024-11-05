package database

import (
	"log"

	"github.com/7Chethan007/PESU-IO_GoLang_Final_Project/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Auto-migrate the models
	err = DB.AutoMigrate(&models.User{}, &models.Question{}, &models.TestCase{})
	if err != nil {
		log.Fatalf("failed to migrate database models: %v", err)
	}

	log.Println("Database initialized and models migrated successfully")
}
