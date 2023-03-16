package category

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) Create(ctx *fiber.Ctx) error{
		// Crear un nuevo usuario
		category := new(models.Category)

		if err := ctx.BodyParser(category); err != nil {
			return err
		}
	
		// Crear el review en la base de datos
	
		if result := database.DB.Preload("Books").Create(category); result.Error != nil {
			return result.Error
		}
	
		// Devolver la respuesta JSON con el objeto creado y un estado HTTP 200
		return ctx.Status(fiber.StatusOK).JSON(category)
}