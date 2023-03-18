package user

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	user := new(models.User)

	err := GetUserById(user, id)

	if err != nil {
		return err
	}

	if err := ctx.BodyParser(user); err != nil {
		return err
	}
	
	if err := UpdateUser(user); err != nil {
		return err
	}
	return ctx.JSON(user)
}


func UpdateUser(user *models .User) error {
	result := database.DB.Save(user)
	return result.Error
}