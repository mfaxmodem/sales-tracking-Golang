package models

import (
	"gorm.io/gorm"
	"time"
)

type Route struct {
	gorm.Model
	UserID    uint      `json:"user_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Visits    []Visit   `json:"visits" gorm:"foreignKey:RouteID"`
}
