package book

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) Delete(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	book := new(models.Book)

	err := GetBookById(book, id)
	
	if err != nil {
		return err
	}

	if result := database.DB.Preload("Reviews").Delete(book); result.Error != nil {
		return result.Error
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}