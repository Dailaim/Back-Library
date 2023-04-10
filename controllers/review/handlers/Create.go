package handlers

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

// @Summary Create review
// @Tags Review
// @Description Create review
// @Produce json
// @Param review body models.Review true "Review"
// @Success 200 {object} models.Review
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /crud/review [post]
func Create(ctx *fiber.Ctx) error {

	// Crear un nuevo usuario
	review := new(models.Review)

	err := ctx.BodyParser(review)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.Error{
			Message: "No se pudo crear el review",
			Code:    fiber.StatusBadRequest,
			Error:   err,
		})
	}

	// Crear el review en la base de datos

	if result := database.DB.Create(review); result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Error{
			Message: "No se pudo crear el review",
			Code:    fiber.StatusInternalServerError,
			Error:   result.Error,
		})
	}

	// Devolver la respuesta JSON con el objeto creado y un estado HTTP 200
	return ctx.Status(fiber.StatusOK).JSON(review)
}
