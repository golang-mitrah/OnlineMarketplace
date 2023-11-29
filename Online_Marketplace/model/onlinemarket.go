package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `json:"username" validate:"required" gorm:"column:username;unique;not null"`
	Password string `json:"password" validate:"required" gorm:"column:password;not null"`
	Token    string `json:"token" gorm:"column:token"`
}

type Product struct {
	gorm.Model
	Name        string  `json:"name" gorm:"type:varchar(100);not null"`
	Description string  `json:"description" gorm:"type:varchar(200);not null"`
	Price       float64 `json:"price" gorm:"type:DECIMAL(10, 2);not null"`
}
