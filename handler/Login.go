package handler

import (
	"github.com/gin-gonic/gin"
	"login_register/models"
	"login_register/services"
	"net/http"
)

func Login(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

func LoginHandler(c *gin.Context) {

	var user models.User
	username := c.PostForm("username")
	password := c.PostForm("password")

	//connect to database and search for user
	db := services.ConnectDB()
	result := db.Where(models.User{Username: username, Password: password}).Find(&user)

	//check if username and password are not correct
	if user.Username != username && user.Password != password {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "Username or Password is incorrect"})
		return
	}

	//check if everything is ok ==> log in the user
	if result != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "logged in!"})
		return
	}

}
