package user

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetById(ctx *fiber.Ctx) error{

	id := ctx.Params("id")

	user := new(models.User) 

	err := GetUserById(user, id)

	if err != nil {
		return err
	}
	
	return ctx.JSON(user)
}

func GetUserById (user *models.User, id string) error{
	
	result := database.DB.Preload("Reviews").First(user, id)

	return result.Error
}


