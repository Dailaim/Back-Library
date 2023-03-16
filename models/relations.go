package models

type BookAuthor struct {
    BookID   uint `gorm:"foreignKey:BookID"`
    AuthorID uint `gorm:"foreignKey:AuthorID"`
}

type BookCategory struct {
    BookID     uint `gorm:"foreignKey:BookID"`
    CategoryID uint `gorm:"foreignKey:CategoryID"`
}