package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// ProfileHandle retrieves the profile of the currently authenticated user
func ProfileHandle(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user ID from the JWT or session (for now, let's assume we get it from query param or JWT middleware)
		userID := c.GetString("userID") // This assumes a middleware sets userID in the context

		var user User
		if err := db.First(&user, "id = ?", userID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		})
	}
}

// UsersHandle handles the request to get all users

func UsersHandle(db *gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		var users []User

		if err := db.Find(&users).Error; err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return

		}

		c.JSON(http.StatusOK, users)

	}

}
