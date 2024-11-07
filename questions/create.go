package questions

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Question struct {
	ID       int    `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func CreateQuestionHandle(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) { // returnss a Gin HandlerFunc that ahndles HTTP requests to create questions
		var question Question // Parse JSON request body
		if err := c.ShouldBindJSON(&question); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		// Insert question into database
		_, err := db.Exec("INSERT INTO questions (question, answer) VALUES (?, ?)", question.Question, question.Answer)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create question"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Question created successfully"})
	}
}
