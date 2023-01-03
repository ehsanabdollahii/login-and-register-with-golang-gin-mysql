package handler

import (
	"github.com/gin-gonic/gin"
	"login_register/models"
	"login_register/services"
)

func Signup(c *gin.Context) {
	c.HTML(200, "signup.html", gin.H{})
}

func SignupHandler(c *gin.Context) {
	c.HTML(200, "signup-successfully.html", gin.H{})

	email := c.PostForm("email")
	username := c.PostForm("username")
	password := c.PostForm("password")
	confirm_password := c.PostForm("confirm_password")

	newUser := models.User{
		Email:           email,
		Username:        username,
		Password:        password,
		ConfirmPassword: confirm_password,
	}

	db := services.ConnectDB()

	var user models.User
	result := db.Where(models.User{Email: email}).Find(&user)

	//check if user is already exists
	if result.RowsAffected > 0 {
		c.JSON(200, gin.H{
			"Message": "Email already exists",
		})
		return
	}

	//check if password and confirm password are equal
	if newUser.Password != newUser.ConfirmPassword {
		c.JSON(200, gin.H{
			"Message": "Password and Confirm Password are not equal",
		})
		return
	}

	//if everything is ok ==> create a new user
	if result.RowsAffected == 0 {
		db.Create(&newUser)

		err := db.AutoMigrate(&newUser)
		if err != nil {
			return
		}
	}
}
