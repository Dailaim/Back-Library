package handlers

import (
	/* "fmt" */

	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
	/* "github.com/google/uuid" */)


// @Summary Create books
// @Tags Book
// @Description Create book
// @Produce json
// @Success 200 {object} models.Book
// @Failure 400
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /crud/book [post]
func Create(ctx *fiber.Ctx) error {
	// Crear un nuevo libro
	book := new(models.Book)

	// Analizar el cuerpo del JSON para obtener los datos del libro
	err := ctx.BodyParser(book)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.Error{
			Message: "No se pudo crear el libro",
			Code:    fiber.StatusBadRequest,
			Error:  err,
		})
	}

	// Crear el libro en la base de datos
	result := database.DB.Preload("Authors").Preload("Categories").Create(book)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Error{
			Message: "No se pudo crear el libro",
			Code:    fiber.StatusInternalServerError,
			Error:  result.Error,
		})
	}

	// Devolver la respuesta JSON con el objeto creado y un estado HTTP 200
	return ctx.Status(fiber.StatusOK).JSON(book)
}

