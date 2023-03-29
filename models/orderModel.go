package models

import "gorm.io/gorm"

type Order struct{
	gorm.Model
	ID uint `gorm:"primaryKey"`
	UserID int `json:"user_id"`
	ProductID int `json:"product_id"`
	Quantity int `json:"quantity"`
	Price int `json:"price"`
	Payed bool `json:"payed"`
}