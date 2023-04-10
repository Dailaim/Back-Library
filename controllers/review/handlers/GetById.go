package handlers

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

// @Summary Get review by id
// @Tags Review
// @Description Get review by id
// @Produce json
// @Param id path string true "Review ID"
// @Success 200 {object} models.Review
// @Failure 404 {object} models.Error
// @Router /crud/review/{id} [get]
func GetById(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	review := new(models.Review)

	err := GetReviewById(review, id)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(models.Error{
			Message: "No se encontró la reseña",
			Code:    fiber.StatusNotFound,
			Error:   err,
		})

	}

	return ctx.Status(fiber.StatusOK).JSON(review)
}

func GetReviewById(review *models.Review, id string) error {

	result := database.DB.First(review, id)

	return result.Error
}
