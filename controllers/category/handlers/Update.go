package handlers

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)


// @Summary Update category
// @Tags Category
// @Description Update category
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} models.Category
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /crud/categories/{id} [put]
func Update(ctx *fiber.Ctx) error {
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

	if err := ctx.BodyParser(category); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.Error{
			Message: "Error al procesar la petición",
			Code:    fiber.StatusBadRequest,
			Error:  err,
		})
	}
	
	if err := UpdateCategory(category); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Error{
			Message: "Error al actualizar la categoría",
			Code:    fiber.StatusInternalServerError,
			Error:  err,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(category)
}

func UpdateCategory(review *models.Category) error {
	result := database.DB.Save(review)
	return result.Error
}