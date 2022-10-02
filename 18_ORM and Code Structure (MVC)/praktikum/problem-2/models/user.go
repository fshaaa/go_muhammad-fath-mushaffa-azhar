package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" from:"name"`
	Email    string `json:"email" from:"email"`
	Password string `json:"password" from:"password"`
}
