package questions

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FetchQuestionsHandle(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT id, question, answer FROM questions")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch questions"})
			return
		}
		defer rows.Close()

		var questions []Question
		for rows.Next() {
			var question Question
			if err := rows.Scan(&question.ID, &question.Question, &question.Answer); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan question"})
				return
			}
			questions = append(questions, question)
		}

		c.JSON(http.StatusOK, questions)
	}
}
