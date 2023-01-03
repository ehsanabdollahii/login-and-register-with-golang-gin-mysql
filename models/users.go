package models

type User struct {
	Email           string `json:"email" form:"email"`
	Username        string `json:"username" form:"username"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}
