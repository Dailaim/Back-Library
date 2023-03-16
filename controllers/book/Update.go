package book

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	book := new(models.Book)

	err := GetBookById(book, id)

	if err != nil {
		return err
	}

	if err := ctx.BodyParser(book); err != nil {
		return err
	}
	
	if err := UpdateBook(book); err != nil {
		return err
	}
	return ctx.JSON(book)
}


func UpdateBook(book *models.Book) error {
	result := database.DB.Save(book)
	return result.Error
}