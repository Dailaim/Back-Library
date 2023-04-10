package handlers

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

// @Summary Delete category
// @Tags Category
// @Description Delete category
// @Produce json
// @Param id path string true "Category ID"
// @Success 204
// @Failure 400
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /crud/categories/{id} [delete]
func Delete(ctx *fiber.Ctx) error {

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

	if result := database.DB.Delete(category); result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Error{
			Message: "No se pudo eliminar la categoría",
			Code:    fiber.StatusInternalServerError,
			Error:  result.Error,
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}