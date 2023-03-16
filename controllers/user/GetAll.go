package user

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetAll(ctx *fiber.Ctx) error {

	var users []*models.User

	result := database.DB.Find(&users)
	if result.Error != nil {
		return result.Error
	}

	for _, user := range users{
		user.Password = ""
	} 
	
	return ctx.Status(200).JSON(users)
}
