package book

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetAll(ctx *fiber.Ctx) error {
	book := new([]models.Book)

	result := database.DB.Preload("Authors").Preload("Categories").Preload("Reviews").Find(book)
	if result.Error != nil {
		return result.Error
	}
	
	return ctx.Status(200).JSON(book)
}
