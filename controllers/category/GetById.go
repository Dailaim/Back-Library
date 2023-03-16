package category

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetById(ctx *fiber.Ctx) error{

	id := ctx.Params("id")

	category := new(models.Category) 

	err := GetCategoryById(category, id)

	if err != nil {
		return err
	}
	
	return ctx.JSON(category)
}

func GetCategoryById (category *models.Category, id string) error{
	
	result := database.DB.Preload("Books").First(category, id)

	return result.Error
}

func GetCategoriesById (category *models.Category, ids []string) error{
	
	result := database.DB.Find(category, ids)

	return result.Error
}
