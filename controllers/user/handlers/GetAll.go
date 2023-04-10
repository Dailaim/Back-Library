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
func GetAll(ctx *fiber.Ctx) error {

	var users []*models.User

	result := database.DB.Find(&users)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Error{
			Message: "No se pudieron obtener los usuarios",
			Code:    fiber.StatusInternalServerError,
			Error:   result.Error,
		})
	}

	for _, user := range users {
		user.Password = ""
	}

	return ctx.Status(fiber.StatusOK).JSON(users)
}
