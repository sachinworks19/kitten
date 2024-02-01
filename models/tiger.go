package models

import (
	"time"

	"gorm.io/gorm"
)

type Tiger struct {
	gorm.Model

	TigerName  string     `json:"tigerName"`
	DOB        time.Time  `json:"dob"`
	LastSeen   time.Time  `json:"lastseen"`
	LastCoords []Location `json:"locations" gorm:"locations"`
	IsActive   bool       `json:"isActive"`
}

type Location struct {
	gorm.Model

	Lat  float64 `gorm:"type:decimal(10,8)"`
	Long float64 `gorm:"type:decimal(10,8)"`
}
