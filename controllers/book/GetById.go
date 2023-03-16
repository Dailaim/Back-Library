package book

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetById(ctx *fiber.Ctx) error{
	id := ctx.Params("id")

	book := new(models.Book) 

	err := GetBookById(book, id)

	if err != nil {
		return err
	}
	
	return ctx.JSON(book)
}

func GetBookById (book *models.Book, id string) error{
	
	result := database.DB.Preload("Authors").Preload("Categories").Preload("Reviews").First(book, id)

	return result.Error
}
