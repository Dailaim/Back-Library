package category

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	category := new(models.Category)

	err := GetCategoryById(category, id)

	if err != nil {
		return err
	}

	if err := ctx.BodyParser(category); err != nil {
		return err
	}
	
	if err := UpdateCategory(category); err != nil {
		return err
	}
	return ctx.JSON(category)
}

func UpdateCategory(review *models.Category) error {
	result := database.DB.Save(review)
	return result.Error
}