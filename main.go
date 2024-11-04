package main

import (
	"github.com/7Chethan007/PESU-IO_GoLang_Final_Project/auth"
	"github.com/7Chethan007/PESU-IO_GoLang_Final_Project/compiler"
	"github.com/7Chethan007/PESU-IO_GoLang_Final_Project/questions"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/auth/signin", auth.Signin)
	router.POST("/auth/signup", auth.Signup)

	router.POST("/run", compiler.Run)

	router.POST("/question/create", questions.CreateQuestion)
	router.POST("/question/fetch", questions.FetchQuestion)
	router.Run(":6969")
}
