package models

type Author struct {
	ID        uint    `gorm:"primaryKey"`
	FirstName string  
	LastName  string  
	Age       int     
	Books     []*Book `gorm:"many2many:book_authors;"`
}
