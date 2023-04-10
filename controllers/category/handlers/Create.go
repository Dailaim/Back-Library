package handlers

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

// @Summary Create category
// @Tags Category
// @Description Create category
// @Produce json
// @Success 200 {object} models.Category
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /crud/categories [post]
func Create(ctx *fiber.Ctx) error{
		// Crear un nuevo usuario
		category := new(models.Category)

		if err := ctx.BodyParser(category); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(models.Error{
				Message: "No se pudo crear la categoria",
				Code:    fiber.StatusBadRequest,
				Error:  err,
			})
		}
	
		// Crear el review en la base de datos
	
		if result := database.DB.Preload("Books").Create(category); result.Error != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(models.Error{
				Message: "No se pudo crear la categoria",
				Code:    fiber.StatusInternalServerError,
				Error:  result.Error,
			})
		}
	
		// Devolver la respuesta JSON con el objeto creado y un estado HTTP 200
		return ctx.Status(fiber.StatusOK).JSON(category)
}