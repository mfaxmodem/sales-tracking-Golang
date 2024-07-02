package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string  `json:"name"`
	Email    string  `json:"email" gorm:"unique"`
	Password string  `json:"password"`
	Role     string  `json:"role"`
	Visits   []Visit `json:"visits" gorm:"foreignKey:VisitorID"`
}
