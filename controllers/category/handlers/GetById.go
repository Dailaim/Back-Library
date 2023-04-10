package handlers

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

// @Summary Get category by id
// @Tags Category
// @Description Get category by id
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} models.Category
// @Failure 400
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /crud/categories/{id} [get]
func GetById(ctx *fiber.Ctx) error{

	id := ctx.Params("id")

	category := new(models.Category) 

	err := GetCategoryById(category, id)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(models.Error{
			Message: "No se encontró la categoría",
			Code:    fiber.StatusNotFound,
			Error:  err,
		})
	}
	
	return ctx.Status(fiber.StatusOK).JSON(category)
}

func GetCategoryById (category *models.Category, id string) error{
	
	result := database.DB.Preload("Books").First(category, id)

	return result.Error
}

func GetCategoriesById (category *models.Category, ids []string) error{
	
	result := database.DB.Find(category, ids)

	return result.Error
}
