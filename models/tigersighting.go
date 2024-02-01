package models

import "gorm.io/gorm"

type TigerSighting struct {
	gorm.Model

	TigerID    string     `json:"tigerID"`
	Location   []Location `json:"locations" gorm:"locations"`
	TigerImage string     `json:"tigerImage"`
}
