package review

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetById(ctx *fiber.Ctx) error{

	id := ctx.Params("id")

	review := new(models.Review) 

	err := GetReviewById(review, id)

	if err != nil {
		return err
	}
	
	return ctx.JSON(review)
}

func GetReviewById (review *models.Review, id string) error{
	
	result := database.DB.First(review, id)

	return result.Error
}


