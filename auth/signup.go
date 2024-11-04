package auth

import (
	"database/sql"
	"net/http"

	"github.com/7Chethan007/PESU-IO_GoLang_Final_Project/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {

	var request models.SignUpRequest
	err := c.BindJSON(&request)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid input",
		})
	}

	// implement
	c.JSON(200, gin.H{
		"success": true,
	})
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// SignupHandler registers a new user with a hashed password
func SignupHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// Hash password before storing
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)

		query := `INSERT INTO users (username, password) VALUES (?, ?)`
		_, err := db.Exec(query, user.Username, hashedPassword)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Signup successful"})
	}
}
