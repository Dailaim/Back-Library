package handlers

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

// @Summary Update book
// @Tags Book
// @Description Update book
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} models.Book
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /crud/books/{id} [put]
func Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	book := new(models.Book)

	err := GetBookById(book, id)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(models.Error{
			Message: "No se encontró el libro",
			Code:    fiber.StatusNotFound,
			Error:  err,
		})
	}

	if err := ctx.BodyParser(book); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.Error{
			Message: "No se pudo actualizar el libro",
			Code:    fiber.StatusBadRequest,
			Error:  err,
		})
	}
	
	if err := UpdateBook(book); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.Error{
			Message: "No se pudo actualizar el libro",
			Code:    fiber.StatusBadRequest,
			Error:  err,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(book)
}


func UpdateBook(book *models.Book) error {
	result := database.DB.Save(book)
	return result.Error
}