package author

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetAll(ctx *fiber.Ctx) error {
	author := new([]models.Author)

	result := database.DB.Find(author)
	if result.Error != nil {
		return result.Error
	}
	
	return ctx.Status(200).JSON(author)
}
