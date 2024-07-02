package models

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Visits    []Visit `json:"visits" gorm:"foreignKey:StoreID"`
}
