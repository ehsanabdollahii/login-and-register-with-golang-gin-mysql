package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "gorm.io/driver/mysql"
	"login_register/handler"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./static/*.html")
	router.Static("/static", "./static/")
	router.GET("/login", handler.Login)
	router.POST("/login", handler.LoginHandler)
	router.GET("/signup", handler.Signup)
	router.POST("/signup", handler.SignupHandler)

	router.Run(":8068")
}
