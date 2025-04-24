package models

import (
	"time"
)

type PersonalAccessToken struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	UserID     uint       `gorm:"index;not null" json:"user_id"`
	Name       string     `gorm:"size:255;not null" json:"name"`
	TokenHash  string     `gorm:"size:255;not null" json:"-"`
	Abilities  string     `gorm:"type:text" json:"abilities"`
	LastUsedAt *time.Time `gorm:"default:NULL" json:"last_used_at,omitempty"`
	ExpiresAt  *time.Time `gorm:"default:NULL" json:"expires_at,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

// TableName is Database TableName of this model
func (e *PersonalAccessToken) TableName() string {
	return "personal_access_token"
}
