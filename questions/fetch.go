package questions

import (
	"database/sql"
	"net/http"

	"github.com/7Chethan007/PESU-IO_GoLang_Final_Project/models"
	"github.com/gin-gonic/gin"
)

func FetchQuestion(c *gin.Context) {
	var request models.FetchQuestionRequest
	c.BindJSON(&request)

	var question models.Question
	c.JSON(200, gin.H{
		"success":  true,
		"question": question,
	})

}

// FetchQuestionsHandler retrieves a list of questions with optional filters
func FetchQuestionsHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		difficulty := c.Query("difficulty")
		query := "SELECT id, title, content, difficulty FROM questions WHERE 1=1"

		if difficulty != "" {
			query += " AND difficulty = ?"
		}

		rows, err := db.Query(query, difficulty)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch questions"})
			return
		}
		defer rows.Close()

		var questions []Question
		for rows.Next() {
			var q Question
			rows.Scan(&q.ID, &q.Title, &q.Content, &q.Difficulty)
			questions = append(questions, q)
		}

		c.JSON(http.StatusOK, gin.H{"questions": questions})
	}
}
