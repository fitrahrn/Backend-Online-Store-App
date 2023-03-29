package models

import "gorm.io/gorm"

type Account struct{
	gorm.Model
	ID uint `gorm:"primaryKey"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}