package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/7Chethan007/PESU-IO_GoLang_Final_Project/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestSignUp(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	db.AutoMigrate(&models.User{})

	router := gin.Default()
	router.POST("/signup", SignupHandle(db))

	t.Run("valid signup", func(t *testing.T) {
		user := models.SignUpRequest{
			Username: "testuser",
			Password: "password123",
		}
		body, _ := json.Marshal(user)
		req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}

		var response map[string]string
		json.Unmarshal(w.Body.Bytes(), &response)
		if response["message"] != "Signup successful" {
			t.Errorf("Expected message 'Signup successful', got '%s'", response["message"])
		}
	})

	t.Run("invalid signup", func(t *testing.T) {
		user := map[string]string{
			"username": "testuser",
			// Missing password field
		}
		body, _ := json.Marshal(user)
		req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
		}

		var response map[string]string
		json.Unmarshal(w.Body.Bytes(), &response)
		if response["error"] != "Invalid request" {
			t.Errorf("Expected error 'Invalid request', got '%s'", response["error"])
		}
	})

	t.Run("signup with existing username", func(t *testing.T) {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		db.Create(&models.User{Username: "existinguser", Password: string(hashedPassword)})

		user := models.SignUpRequest{
			Username: "existinguser",
			Password: "password123",
		}
		body, _ := json.Marshal(user)
		req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusInternalServerError {
			t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, w.Code)
		}

		var response map[string]string
		json.Unmarshal(w.Body.Bytes(), &response)
		if response["error"] != "Failed to create user" {
			t.Errorf("Expected error 'Failed to create user', got '%s'", response["error"])
		}
	})
}
