package handlers

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

// @Summary Get user by email
// @Tags User
// @Description Get user by email
// @Produce json
// @Param email path string true "Email of the user"
// @Success 200 {object} models.User
// @Failure 404 {object} models.Error
// @Router /crud/user [get]
func GetByEmail(ctx *fiber.Ctx) error {
	user := new(models.User)

	err := ctx.BodyParser(user)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.Error{
			Message: "No se pudo obtener el usuario",
			Code:    fiber.StatusBadRequest,
			Error:   err,
		})
	}
	
	err = GetUserByEmail(user, user.Email)
	
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(models.Error{
			Message: "No se encontró el usuario",
			Code:    fiber.StatusNotFound,
			Error:   err,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(user)
}

func GetUserByEmail(user *models.User, email string) error {
	result := database.DB.Where("email = ?", email).First(user)
	return result.Error
}