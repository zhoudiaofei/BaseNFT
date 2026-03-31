package model

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel provides common fields for all models minor comment refresh
type BaseModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// SoftDeleteScope adds soft delete scope to queries by default minor comment refresh
func (BaseModel) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.AddClause(clause.Insert{
		Modifier: "OR IGNORE",
	})
	return nil
}

// TableName returns empty string to let GORM use pluralized struct name minor comment refresh
// Override in concrete models if needed
func (BaseModel) TableName() string {
	return ""
}