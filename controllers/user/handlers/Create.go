package handlers

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

func Create(ctx *fiber.Ctx) error {

	// Crear un nuevo usuario
	user := new(models.User)

	err := ctx.BodyParser(user)

	if err != nil {
		return err
	}

	err = CreateUser(user)

	if err != nil {
		return err
	}

	// Devolver la respuesta JSON con el objeto creado y un estado HTTP 200
	return ctx.Status(fiber.StatusOK).JSON(user)
}

func CreateUser(user *models.User) error {

	// Crear el usuario en la base de datos

	result := database.DB.Create(user)

	return result.Error

}
