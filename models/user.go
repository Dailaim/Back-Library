package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"not null;unique_index"`
	Name     string `json:"name" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Photo    string `json:"photo"  gorm:"not null"`
}
