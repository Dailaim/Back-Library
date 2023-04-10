package handlers

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

func Update(ctx *fiber.Ctx) error {
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

	if err := ctx.BodyParser(review); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.Error{
			Message: "No se pudo actualizar la reseña",
			Code:    fiber.StatusBadRequest,
			Error:   err,
		})
	}

	if err := UpdateReview(review); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Error{
			Message: "No se pudo actualizar la reseña",
			Code:    fiber.StatusInternalServerError,
			Error:   err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(review)
}

func UpdateReview(review *models.Review) error {
	result := database.DB.Save(review)
	return result.Error
}
