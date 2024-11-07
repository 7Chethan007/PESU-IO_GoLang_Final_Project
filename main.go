package main

import (
	"log"

	"github.com/7Chethan007/PESU-IO_GoLang_Final_Project/auth"
	"github.com/7Chethan007/PESU-IO_GoLang_Final_Project/auth/middleware"
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

	// Public routes
	router.POST("/auth/signin", auth.SigninHandle(database.DB))
	router.POST("/auth/signup", auth.SignupHandle(database.DB))

	// Protected routes
	protected := router.Group("/auth")
	protected.Use(middleware.AuthMiddleware()) // Apply JWT middleware

	// Sub-Part-Of-Protected routes
	// Fetches and returns the profile information of the authenicated user
	protected.GET("/profile", auth.ProfileHandle(database.DB)) // Profile route	- passing DB connection in parameter
	// Fetches and returns all the users in the system
	protected.GET("/users", auth.UsersHandle(database.DB)) // View all users route - passing DB connection in parameter

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
