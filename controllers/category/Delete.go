package category

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) Delete(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	category := new(models.Category)

	err := GetCategoryById(category, id)
	
	if err != nil {
		return err
	}

	if result := database.DB.Delete(category); result.Error != nil {
		return result.Error
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}