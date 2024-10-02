package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code        string `json:"code"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}
