package handlers

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)
// @Summary Get book by id
// @Tags Book
// @Description Get book by id
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} models.Book
// @Failure 400
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /crud/book/{id} [get]
func GetById(ctx *fiber.Ctx) error{
	id := ctx.Params("id")

	book := new(models.Book) 

	err := GetBookById(book, id)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(models.Error{
			Message: "No se encontró el libro",
			Code:    fiber.StatusNotFound,
		})
	}
	
	return ctx.Status(fiber.StatusOK).JSON(book)
}

func GetBookById (book *models.Book, id string) error{
	
	result := database.DB.Preload("Authors").Preload("Categories").Preload("Reviews").First(book, id)

	return result.Error
}
