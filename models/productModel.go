package models

import "gorm.io/gorm"

type Product struct{
	gorm.Model
	ID uint `gorm:"primaryKey"`
	Name string `json:"name"`
	Description string `json:"description"`
	Quantity int `json:"quantity"`
	Price int `json:"price"`
	Seller string `json:"seller"`
	Category string `json:"category"`

}