package author

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetById(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	author := new(models.Author)

	err := GetAuthorById(author, id)

	if err != nil {
		return err
	}

	return ctx.JSON(author)
}

func GetAuthorById(author *models.Author, id string) error {

	result := database.DB.First(author, id)

	return result.Error
}
