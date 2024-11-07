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
		// Get user ID from the JWT or session
		userID := c.GetString("userID") // Retreving the userID from the request context in postman

		var user User // Interates through the db with userID, if not found returns 404
		if err := db.First(&user, "id = ?", userID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		// On successful iteration of db, returns details of user
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

		var users []User // In user variable we are storing all the users in the array

		if err := db.Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return

		}
		// if no error, returns all the users stored in the array users above
		c.JSON(http.StatusOK, users)

	}

}
