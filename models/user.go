package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginRequest struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}