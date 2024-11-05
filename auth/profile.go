package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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
