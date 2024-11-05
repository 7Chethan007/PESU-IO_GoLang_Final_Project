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

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.POST("/signin", SigninHandle(db))
	return r
}

func TestSignIn(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.User{})

	password, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	db.Create(&models.User{Username: "testuser", Password: string(password)})

	r := setupRouter(db)

	t.Run("valid credentials", func(t *testing.T) {
		body := models.SignInRequest{
			Username: "testuser",
			Password: "password123",
		}
		jsonValue, _ := json.Marshal(body)
		req, _ := http.NewRequest("POST", "/signin", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}
	})

	t.Run("invalid credentials", func(t *testing.T) {
		body := models.SignInRequest{
			Username: "testuser",
			Password: "wrongpassword",
		}
		jsonValue, _ := json.Marshal(body)
		req, _ := http.NewRequest("POST", "/signin", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusUnauthorized {
			t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, w.Code)
		}
	})

	t.Run("user not found", func(t *testing.T) {
		body := models.SignInRequest{
			Username: "nonexistentuser",
			Password: "password123",
		}
		jsonValue, _ := json.Marshal(body)
		req, _ := http.NewRequest("POST", "/signin", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusUnauthorized {
			t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, w.Code)
		}
	})

	t.Run("invalid request body", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/signin", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
		}
	})
}
