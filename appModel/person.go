package appModel

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	isActive string `json:"isActive"`
	Token    string `json:"token"`
}
