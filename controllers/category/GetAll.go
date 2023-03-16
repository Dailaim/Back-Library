package category

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetAll(ctx *fiber.Ctx) error {
	category := new([]models.Category)

	result := database.DB.Preload("Books").Find(category)
	if result.Error != nil {
		return result.Error
	}
	
	return ctx.Status(200).JSON(category)
}
