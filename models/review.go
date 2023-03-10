package models

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserID  uint    `gorm:"user_id"`
    User    User    `gorm:"foreignKey:UserID" json:"user"`
	Comment string `json:"comment"`
	BookID  uint   `json:"book_id"`
	Book    Book   `gorm:"foreignKey:BookID" json:"Book"`
	Score   uint   `json:"score"`
}
