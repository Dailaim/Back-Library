package handlers

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

// @Summary Get all authors
// @Tags Author
// @Description Get all authors
// @Produce json
// @Success 200 {array} []models.Author
// @Failure 500 {object} models.Error
// @Router /crud/author [get]
func GetAll(ctx *fiber.Ctx) error {
	author := new([]models.Author)

	result := database.DB.Find(author)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Error{
			Message: "No se pudieron obtener los autores",
			Code:    fiber.StatusInternalServerError,
			Error: result.Error,
		})
	}
	
	return ctx.Status(200).JSON(author)
}
