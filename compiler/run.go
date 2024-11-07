package compiler

import (
	"database/sql"
	"fmt"
	"os/exec"
	"strings"

	"github.com/7Chethan007/PESU-IO_GoLang_Final_Project/models"
	"github.com/gin-gonic/gin"
)

func Run(c *gin.Context) {
	var request models.RunRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// TODO: Implement logic to execute code based on `request.Language` and `request.Code`.
	// Capture stdout and stderr, then store the result in the database for caching.

	c.JSON(200, gin.H{
		"output": "Hello World", // Placeholder; replace with actual output
	})
}

// ExecuteCode executes code in the specified language with given input
func ExecuteCode(language, code, input string) (string, error) {
	var cmd *exec.Cmd

	switch language {
	case "python":
		cmd = exec.Command("python3", "-c", code)
	case "javascript":
		cmd = exec.Command("node", "-e", code)
	default:
		return "", fmt.Errorf("unsupported language")
	}

	cmd.Stdin = strings.NewReader(input)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

// SaveExecutionResult saves the execution result to the database for caching
func SaveExecutionResult(db *sql.DB, language, code, input, output string) {
	query := `INSERT INTO code_executions (language, code, input, output) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, language, code, input, output)
	if err != nil {
		fmt.Println("Error saving execution result:", err)
	}

	// TODO: Implement logic to save the execution result to the database

}
