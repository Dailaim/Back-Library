package handlers

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

// @Summary Delete user
// @Tags User
// @Description Delete user
// @Produce json
// @Param id path string true "User ID"
// @Success 204
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /crud/user/{id} [delete]
func Delete(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	user := new(models.User)

	err := GetUserById(user, id)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(models.Error{
			Message: "No se encontró el usuario",
			Code:    fiber.StatusNotFound,
			Error:   err,
		})
	}

	if result := database.DB.Delete(user); result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Error{
			Message: "No se pudo eliminar el usuario",
			Code:    fiber.StatusInternalServerError,
			Error:   result.Error,
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
