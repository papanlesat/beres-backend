package models

import "gorm.io/gorm"

type Setting struct {
	gorm.Model
	Key   string `gorm:"size:50;uniqueIndex"`
	Value string `gorm:"type:text"`
}
