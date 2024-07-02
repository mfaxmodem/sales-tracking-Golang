package models

import (
	"gorm.io/gorm"
	"time"
)

type Checkpoint struct {
	gorm.Model
	RouteID   uint      `json:"route_id"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Timestamp time.Time `json:"timestamp"`
}
