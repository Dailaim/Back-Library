package handlers

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

// @Summary Delete review
// @Tags Review
// @Description Delete review
// @Produce json
// @Param id path string true "Review ID"
// @Success 204
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /crud/review/{id} [delete]
func Delete(ctx *fiber.Ctx) error {

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

	if result := database.DB.Delete(review); result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Error{
			Message: "No se pudo eliminar la reseña",
			Code:    fiber.StatusInternalServerError,
			Error:   result.Error,
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)

}
