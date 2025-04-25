package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name        string `gorm:"size:50"`
	Slug        string `gorm:"size:50;uniqueIndex"`
	Description string `gorm:"size:500"`
	Posts       []Post `gorm:"many2many:posts_tags;"`
}
