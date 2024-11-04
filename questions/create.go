package questions

import (
	"database/sql"
	"net/http"

	"github.com/7Chethan007/PESU-IO_GoLang_Final_Project/models"
	"github.com/gin-gonic/gin"
)

func CreateQuestion(c *gin.Context) {
	var request models.CreateQuestionRequest
	c.BindJSON(&request)

	c.JSON(200, gin.H{
		"success":    true,
		"questionID": 1,
	})

}

type Question struct {
	ID         int      `json:"id"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Difficulty string   `json:"difficulty"`
	Tags       []string `json:"tags"`
}

// CreateQuestionHandler handles creating a new question
func CreateQuestionHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var q Question
		if err := c.BindJSON(&q); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// Insert question into database
		query := `INSERT INTO questions (title, content, difficulty) VALUES (?, ?, ?)`
		res, err := db.Exec(query, q.Title, q.Content, q.Difficulty)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create question"})
			return
		}

		// Add tags to database (assuming a simple join table structure)
		questionID, _ := res.LastInsertId()
		for _, tag := range q.Tags {
			db.Exec("INSERT INTO question_tags (question_id, tag) VALUES (?, ?)", questionID, tag)
		}

		c.JSON(http.StatusOK, gin.H{"message": "Question created"})
	}
}
