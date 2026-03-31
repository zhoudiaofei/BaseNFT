package model

import (
	"time"

	"gorm.io/gorm"
)

// Tag represents a content tag entity. minor comment refresh
// It maps to the `tags` table in the database. minor comment refresh
// PROMPT-F78CD1-000080: lightweight, RESTful-ready, with minimal required fields. minor comment refresh
type Tag struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:64;not null;uniqueIndex" json:"name"`
	Slug      string    `gorm:"size:96;not null;uniqueIndex" json:"slug"`
	Color     string    `gorm:"size:16;default:'#3b82f6'" json:"color,omitempty"` // e.g., hex or tailwind class
	IsPublic  bool      `gorm:"default:true" json:"is_public"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BeforeCreate implements GORM hook to auto-generate slug from name if empty.
func (t *Tag) BeforeCreate(tx *gorm.DB) error {
	if t.Slug == "" && t.Name != "" {
		t.Slug = generateSlug(t.Name)
	}
	return nil
}

// generateSlug is a simple ASCII-only slug generator for demo & local dev.
// In production, use a robust library like github.com/gosimple/slug.
func generateSlug(s string) string {
	var result []rune
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '-' || r == '_' {
			result = append(result, r)
		} else if r == ' ' || r == '.' || r == ',' || r == ':' {
			if len(result) == 0 || result[len(result)-1] != '-' {
				result = append(result, '-')
			}
		}
	}
	// Trim leading/trailing dashes
	for len(result) > 0 && (result[0] == '-' || result[len(result)-1] == '-') {
		if result[0] == '-' {
			result = result[1:]
		}
		if len(result) > 0 && result[len(result)-1] == '-' {
			result = result[:len(result)-1]
		}
	}
	return string(result)
}