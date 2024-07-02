package models

import (
	"gorm.io/gorm"
	"time"
)

type Visit struct {
	gorm.Model
	RouteID   uint      `json:"route_id"`
	StoreID   uint      `json:"store_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	VisitorID uint      `json:"visitor_id"`
}
