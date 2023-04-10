package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
)

// @Summary Get user by id
// @Tags User
// @Description Get user by id
// @Produce json
// @Param id path string true "User id"
// @Success 200 {object} models.User
// @Failure 404 {object} models.Error
// @Router /crud/user/{id} [get]
func GetById(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	user := new(models.User)

	err := GetUserById(user, id)

	if err != nil {
		return err
	}
	user.Password = ""

	return ctx.JSON(user)
}

func GetUserById(user *models.User, id string) error {

	result := database.DB.Preload("Reviews").First(user, id)

	return result.Error
}
