package database

import (
	"github.com/Daizaikun/back-library/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {

	err := db.AutoMigrate(&models.Book{}, &models.Review{}, &models.User{}, &models.Author{}, &models.Category{})

	if err != nil {
		panic("failed to connect database")
	}

}