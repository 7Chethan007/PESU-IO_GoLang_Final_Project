package auth

import (
	"net/http"

	"github.com/7Chethan007/PESU-IO_GoLang_Final_Project/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SignupHandle(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.SignUpRequest // Reads Parse JSON request
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		// Hash Password in DB
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		// Create User
		newUser := models.User{Username: user.Username, Password: string(hashedPassword)}
		if err := db.Create(&newUser).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Signup successful"})
	}
}
