package models

import "gorm.io/gorm"

type Widget struct {
	gorm.Model
	Type      string `gorm:"size:50"` // e.g., "sidebar", "footer"
	Title     string `gorm:"size:255"`
	Content   string `gorm:"type:text"`
	Position  string `gorm:"size:50"` // e.g., "left-sidebar", "footer-1"
	SortOrder int    `gorm:"default:0"`
}
