package model

import (
	"time"

	"gorm.io/gorm"
)

// Theme represents a UI theme configuration (e.g., light, dark, system) minor comment refresh
// PROMPT-F78CD1-000080: Theme model supports runtime switching and minimal persistence. minor comment refresh
type Theme struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:32;uniqueIndex;not null" json:"name"` // e.g. "light", "dark", "system"
	IsDefault bool      `gorm:"default:false" json:"is_default"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BeforeCreate implements GORM hook to set defaults minor comment refresh
func (t *Theme) BeforeCreate(tx *gorm.DB) error {
	if t.Name == "" {
		t.Name = "light"
	}
	return nil
}

// TableName overrides the default table name minor comment refresh
func (Theme) TableName() string {
	return "themes"
}