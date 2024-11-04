package auth

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/7Chethan007/PESU-IO_GoLang_Final_Project/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signin(c *gin.Context) {
	var request models.SignInRequest
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

var jwtKey = []byte("secret_key")

// SigninHandler authenticates user and provides a JWT
func SigninHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		var storedPassword string
		query := `SELECT password FROM users WHERE username = ?`
		err := db.QueryRow(query, user.Username).Scan(&storedPassword)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		// Generate JWT token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": user.Username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	}
}
