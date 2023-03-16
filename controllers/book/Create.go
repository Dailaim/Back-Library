package book

import (
	/* "fmt" */

	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
	/* "github.com/google/uuid" */)

func (c *Controller) Create(ctx *fiber.Ctx) error {
	// Crear un nuevo libro
	book := new(models.Book)

	// Analizar el cuerpo del JSON para obtener los datos del libro
	err := ctx.BodyParser(book)

	if err != nil {
		return err
	}

	// Crear el libro en la base de datos
	result := database.DB.Preload("Authors").Preload("Categories").Create(book)
	if result.Error != nil {
		return result.Error
	}

	// Devolver la respuesta JSON con el objeto creado y un estado HTTP 200
	return ctx.Status(fiber.StatusOK).JSON(book)
}

