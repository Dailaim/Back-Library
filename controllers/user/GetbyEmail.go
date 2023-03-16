package user

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

func GetByEmail(ctx *fiber.Ctx) error {
	user := new(models.User)

	err := ctx.BodyParser(user)

	if err != nil {
		return err
	}
	
	err = GetUserByEmail(user, user.Email)
	
	if err != nil {
		return err
	}
	return ctx.JSON(user)
}

func GetUserByEmail(user *models.User, email string) error {
	result := database.DB.Where("email = ?", email).First(user)
	return result.Error
}