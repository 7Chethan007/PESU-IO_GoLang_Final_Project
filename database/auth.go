// package database

// func CreateUser(username, password string) error {
// 	// creates a new user in the database, returns error if any
// 	return nil
// }

//	func CheckPassword(username, password string) (success bool, err error) {
//		// checks if the password is correct for the given username
//		return true, nil
//	}
//
// database/auth.go
package database

import (
	"github.com/7Chethan007/PESU-IO_GoLang_Final_Project/models"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser creates a new user in the database with a hashed password
func CreateUser(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	return DB.Create(&user).Error
}

// CheckPassword verifies if the password is correct for the given username
func CheckPassword(username, password string) (bool, error) {
	var user models.User
	err := DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, nil // Password mismatch
	}

	return true, nil // Password match
}
