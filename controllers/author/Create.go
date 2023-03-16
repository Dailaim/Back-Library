package author

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) Create(ctx *fiber.Ctx) error {
	// Crear un nuevo usuario
	Author := new(models.Author)

	err := ctx.BodyParser(Author)

	if err != nil {
		return err
	}

	// Crear el review en la base de datos

	if result := database.DB.Create(Author); result.Error != nil {
		return result.Error
	}

	// Devolver la respuesta JSON con el objeto creado y un estado HTTP 200
	return ctx.Status(fiber.StatusOK).JSON(Author)
}
