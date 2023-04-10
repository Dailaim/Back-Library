package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
)

// @Summary Get all reviews
// @Tags Review
// @Description Get all reviews
// @Produce json
// @Success 200 {array} []models.Review
// @Failure 500 {object} models.Error
// @Router /crud/review [get]
func GetAll(ctx *fiber.Ctx) error {

	reviews := new([]models.Review)

	result := database.DB.Preload("Book").Preload("User").Find(reviews)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Error{
			Message: "No se pudieron obtener las reseñas",
			Code:    fiber.StatusInternalServerError,
			Error:   result.Error,
		})
	}

	return ctx.Status(200).JSON(reviews)

}
