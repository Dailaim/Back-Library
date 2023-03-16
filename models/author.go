package models

type Author struct {
	ID   uint `gorm:"primaryKey"`
	Name  string  `json:"name,omitempty"`
	Books []*Book `gorm:"many2many:book_authors;" json:"books,omitempty"`
}
