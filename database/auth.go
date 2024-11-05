// package database

// import (
// 	"log"

// 	"github.com/7Chethan007/PESU-IO_GoLang_Final_Project/auth"
// 	"github.com/7Chethan007/PESU-IO_GoLang_Final_Project/compiler"
// 	"github.com/7Chethan007/PESU-IO_GoLang_Final_Project/database"
// 	"github.com/7Chethan007/PESU-IO_GoLang_Final_Project/questions"
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	// Initialize database
// 	database.InitDB()

// 	router := gin.Default()

// 	// Define routes with DB instance passed to handlers
// 	router.POST("/auth/signin", auth.SigninHandle(database.DB))
// 	router.POST("/auth/signup", auth.SignupHandle(database.DB))

// 	router.POST("/run", compiler.Run)

// 	router.POST("/question/create", questions.CreateQuestionHandle(database.DB))
// 	router.GET("/questions", questions.FetchQuestionsHandle(database.DB))

//		// Start server
//		port := ":6969"
//		if err := router.Run(port); err != nil {
//			log.Fatalf("Server failed to start: %v", err)
//		}
//	}
package database
