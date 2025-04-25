package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string     `gorm:"size:50"`
	Slug        string     `gorm:"size:50;uniqueIndex"`
	Description string     `gorm:"size:500"`
	ParentID    *uint      `gorm:"index"`
	Children    []Category `gorm:"foreignKey:ParentID"`
	Posts       []Post     `gorm:"many2many:posts_categories;"`
}
