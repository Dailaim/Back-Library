package models

import (
	"time"

	"gorm.io/gorm"
)

// modelos básico de gorm
type BasicModel struct {
	ID        uint `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt time.Time `json:"create_at,omitempty"`
	UpdatedAt time.Time `json:"Updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}