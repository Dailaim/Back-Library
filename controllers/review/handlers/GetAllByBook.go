package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"

)

// @Summary Get all reviews by book
// @Tags Review
// @Description Get all reviews by book
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {array} []models.Review
// @Failure 500 {object} models.Error
// @Router /crud/review/book/{id} [get]
func GetAllByBook(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	reviews := new([]models.Review)

	err := GetAllReviewByBook(reviews, id)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Error{
			Message: "No se pudieron obtener las reseñas",
			Code:    fiber.StatusInternalServerError,
			Error: err,
		})
	}

	return ctx.Status(200).JSON(reviews)

}	

func GetAllReviewByBook(reviews *[]models.Review, id string) error {

	result := database.DB.Where("book_id = ?", id).Preload("User").Find(reviews)

	return result.Error
}
