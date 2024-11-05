package main

import (
	"log"

	"github.com/7Chethan007/PESU-IO_GoLang_Final_Project/auth"
	"github.com/7Chethan007/PESU-IO_GoLang_Final_Project/compiler"
	"github.com/7Chethan007/PESU-IO_GoLang_Final_Project/database"
	"github.com/7Chethan007/PESU-IO_GoLang_Final_Project/questions"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Initialize database
	database.InitDB()
	sqlDB, err := database.DB.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB from gorm.DB: %v", err)
	}

	// Authentication routes
	router.POST("/auth/signin", auth.SigninHandle(database.DB))
	router.POST("/auth/signup", auth.SignupHandle(database.DB))

	// Add new routes for profile and viewing all users
	router.GET("/auth/profile", auth.ProfileHandle(database.DB)) // For getting user profile
	router.GET("/auth/users", auth.UsersHandle(database.DB))     // For getting all users

	// Compiler and questions routes
	router.POST("/run", compiler.Run)
	router.POST("/question/create", questions.CreateQuestionHandle(sqlDB))
	router.GET("/questions", questions.FetchQuestionsHandle(sqlDB))

	// Start server
	port := "8080"
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
