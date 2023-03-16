package models

type Category struct {
	ID    uint    `gorm:"primaryKey"`
	Name  string  `json:"name,omitempty"`
	Books []*Book `gorm:"many2many:book_categories;" json:"books,omitempty"`
}
