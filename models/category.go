package models

type Category struct {
	ID   uint `gorm:"primaryKey"`
	Name string	`json:"name"`
}
