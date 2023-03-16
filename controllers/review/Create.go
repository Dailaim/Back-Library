package review

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) Create(ctx *fiber.Ctx) error {

	// Crear un nuevo usuario
	review := new(models.Review)

	ctx.BodyParser(review)

	// Crear el review en la base de datos

	if result := database.DB.Create(review); result.Error != nil {
		return result.Error
	}

	// Devolver la respuesta JSON con el objeto creado y un estado HTTP 200
	return ctx.Status(fiber.StatusOK).JSON(review)
}
