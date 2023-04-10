package handlers

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)


// @Summary Get all categories
// @Tags Category
// @Description Get all categories
// @Produce json
// @Success 200 {array} []models.Category
// @Failure 500 {object} models.Error
// @Router /crud/categories [get]
func GetAll(ctx *fiber.Ctx) error {
	category := new([]models.Category)

	result := database.DB.Preload("Books").Find(category)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Error{
			Message: "No se pudieron obtener las categorías",
			Code:    fiber.StatusInternalServerError,
			Error: result.Error,
		})
	}
	
	return ctx.Status(fiber.StatusOK).JSON(category)
}
