// Go ORM Models (using GORM)
package models

import (
	"time"

	"gorm.io/datatypes"
)

type Section struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Name         string         `gorm:"size:255;not null" json:"name"`
	SectionType  string         `gorm:"size:50;not null;index" json:"section_type"`
	DisplayOrder int            `gorm:"default:0;index" json:"display_order"`
	IsActive     bool           `gorm:"default:true;index" json:"is_active"`
	Details      datatypes.JSON `gorm:"type:json;not null" json:"details"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

func (e *Section) TableName() string {
	return "sections"
}

/*
type HeroSectionDetails struct {
	Title      string `json:"title"`
	Subtitle   string `json:"subtitle"`
	ButtonText string `json:"button_text"`
	ButtonLink string `json:"button_link"`
	ImageURL   string `json:"image_url"`
	Alignment  string `json:"alignment" enum:"left,center,right"`
	Overlay    bool   `json:"overlay"`
}

type FeaturesSectionDetails struct {
	Features []FeatureItem `json:"features"`
	Layout   string        `json:"layout" enum:"grid,list,carousel"`
	Columns  int           `json:"columns" enum:"2,3,4"`
}

type FeatureItem struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Link        string `json:"link"`
}
*/
