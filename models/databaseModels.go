package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type Question struct {
	gorm.Model
	Question  string     `json:"question"`
	TestCases []TestCase `json:"testCases" gorm:"foreignKey:QuestionID"` // Define foreign key for TestCases
}

type TestCase struct {
	gorm.Model
	QuestionID     uint   `json:"-"` // Foreign key to associate with Question
	Input          string `json:"input"`
	ExpectedOutput string `json:"expectedOutput"`
}
