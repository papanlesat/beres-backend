package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title         string     `gorm:"size:255"`
	Slug          string     `gorm:"size:255;uniqueIndex"`
	Content       string     `gorm:"type:longtext"`
	Excerpt       string     `gorm:"size:500"`
	AuthorID      uint       `gorm:"index"`
	Author        User       `gorm:"foreignKey:AuthorID"`
	Status        string     `gorm:"size:20;default:'draft';check:status IN ('draft', 'publish', 'trash')"`
	FeaturedImage string     `gorm:"size:255"`
	Categories    []Category `gorm:"many2many:posts_categories;"`
	Tags          []Tag      `gorm:"many2many:posts_tags;"`
}
