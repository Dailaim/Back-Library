package review

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	review := new(models.Review)

	err := GetReviewById(review, id)

	if err != nil {
		return err
	}

	if err := ctx.BodyParser(review); err != nil {
		return err
	}

	if err := UpdateReview(review); err != nil {
		return err
	}

	return ctx.JSON(review)
}

func UpdateReview(review *models.Review) error {
	result := database.DB.Save(review)
	return result.Error
}
