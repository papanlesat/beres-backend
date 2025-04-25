package models

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Name     string `gorm:"size:50"`
	Location string `gorm:"size:50;uniqueIndex"`
	Items    []MenuItem
}

type MenuItem struct {
	gorm.Model
	MenuID   uint       `gorm:"index"`
	ParentID *uint      `gorm:"index"`
	Title    string     `gorm:"size:100"`
	URL      string     `gorm:"size:255"`
	Order    int        `gorm:"default:0"`
	Class    string     `gorm:"size:50"`
	Target   string     `gorm:"size:20"`
	Children []MenuItem `gorm:"foreignKey:ParentID"`
}
