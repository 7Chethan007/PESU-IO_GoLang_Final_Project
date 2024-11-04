package database

import (
	"log"

	"github.com/7Chethan007/PESU-IO_GoLang_Final_Project/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(databaseFileName string) {
	// implement
	// populate DB variable
	var err error
	DB, err = gorm.Open(sqlite.Open(databaseFileName), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Migrate the schema for User, Question, and TestCase
	err = DB.AutoMigrate(&models.User{}, &models.Question{}, &models.TestCase{})
	if err != nil {
		log.Fatalf("failed to migrate database models: %v", err)
	}

	log.Println("Database connection initialized and models migrated successfully")
}
