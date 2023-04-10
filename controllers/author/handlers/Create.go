package handlers

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

// @Summary Create author
// @Tags Author
// @Description Create author
// @Produce json
// @Success 200 {object} models.Author <-- This is a user defined struct.
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /crud/author [post]
func Create(ctx *fiber.Ctx) error {
	// Crear un nuevo usuario
	Author := new(models.Author)

	err := ctx.BodyParser(Author)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.Error{
			Message: "No se pudo crear el autor",
			Code:    fiber.StatusBadRequest,
			Error:  err,
		})
	}

	// Crear el review en la base de datos

	if result := database.DB.Create(Author); result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Error{
			Message: "No se pudo crear el autor",
			Code:    fiber.StatusInternalServerError,
			Error:  result.Error,
		})
	}

	// Devolver la respuesta JSON con el objeto creado y un estado HTTP 200
	return ctx.Status(fiber.StatusOK).JSON(Author)
}
