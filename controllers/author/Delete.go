package author

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) Delete(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	author := new(models.Author)

	err := GetAuthorById(author, id)
	
	if err != nil {
		return err
	}

	if result := database.DB.Delete(author); result.Error != nil {
		return result.Error
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}