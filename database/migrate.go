package database

import (
	"github.com/Daizaikun/back-library/models"
	"gorm.io/gorm"
)

//Verifica y crea las tablas de la base de datos
func Migrate(db *gorm.DB) error{

	err := db.AutoMigrate(&models.Book{}, &models.Review{}, &models.User{},
		&models.Author{}, &models.Category{}, &models.BookAuthor{}, &models.BookCategory{})

	if err != nil {
		return err
	}
	
	return nil
}
