package review

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
)

func (c *Controller) GetAll(ctx *fiber.Ctx) error {

	reviews := new([]models.Review)

	result := database.DB.Preload("Book").Preload("User").Find(reviews)
	if result.Error != nil {
		return result.Error
	}

	return ctx.Status(200).JSON(reviews)

}
