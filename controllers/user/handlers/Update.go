package handlers

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

// @Summary Get all users
// @Tags User
// @Description Get all users
// @Produce json
// @Success 200 {array} []models.User
// @Failure 500 {object} models.Error
// @Router /crud/user [get]
func Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	user := new(models.User)

	err := GetUserById(user, id)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(models.Error{
			Message: "Usuario no encontrado",
			Code:    fiber.StatusNotFound,
			Error:   err,
		})
	}

	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.Error{
			Message: "Error al parsear el body",
			Code:    fiber.StatusBadRequest,
			Error:   err,
		})
	}

	if err := UpdateUser(user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Error{
			Message: "Error al actualizar el usuario",
			Code:    fiber.StatusInternalServerError,
			Error:   err,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(user)
}

func UpdateUser(user *models.User) error {
	result := database.DB.Save(user)
	return result.Error
}
