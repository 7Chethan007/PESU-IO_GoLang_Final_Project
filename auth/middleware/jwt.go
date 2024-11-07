package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("my_test_secret_key") // Replace this with a secure secret

// GenerateJWT generates a JWT token for a given user ID
func GenerateJWT(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 72).Unix(), // Token expiration, e.g., 72 hours
	})

	return token.SignedString(jwtSecret)
}

// ParseJWT parses and validates a JWT token
func ParseJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["userID"].(string)
		return userID, nil
	}

	return "", jwt.ErrSignatureInvalid
}
