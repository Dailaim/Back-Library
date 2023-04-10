package handlers

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

// @Summary Get all books
// @Tags Book
// @Description Get all books
// @Produce json
// @Success 200 {array} []models.Book
// @Failure 500 {object} models.Error
// @Router /crud/books [get]
func GetAll(ctx *fiber.Ctx) error {
	book := new([]models.Book)

	result := database.DB.Preload("Authors").Preload("Categories").Find(book)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Error{
			Message: "No se pudieron obtener los libros",
			Code:    fiber.StatusInternalServerError,
		})
	}
	
	return ctx.Status(fiber.StatusOK).JSON(book)
}
