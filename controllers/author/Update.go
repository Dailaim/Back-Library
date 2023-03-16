package author

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	author := new(models.Author)

	err := GetAuthorById(author, id)

	if err != nil {
		return err
	}

	if err := ctx.BodyParser(author); err != nil {
		return err
	}
	
	if err := UpdateAuthor(author); err != nil {
		return err
	}
	return ctx.JSON(author)
}


func UpdateAuthor(author *models.Author) error {
	result := database.DB.Save(author)
	return result.Error
}
