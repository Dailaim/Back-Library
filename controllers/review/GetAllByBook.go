package review

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllByBook(ctx *fiber.Ctx) error {
	
	id := ctx.Params("id")

	reviews := new([]models.Review)

	err := GetAllReviewByBook(reviews, id)

	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(reviews)
	
}

func GetAllReviewByBook (reviews *[]models.Review, id string) error{
	
	result := database.DB.Where("book_id = ?", id).Find(reviews)

	return result.Error
}
