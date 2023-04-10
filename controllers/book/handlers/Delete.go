package handlers

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)
// @Summary Delete book
// @Tags Book
// @Description Delete book
// @Produce json
// @Param id path string true "Book ID"
// @Success 204
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /crud/books/{id} [delete]
func Delete(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	book := new(models.Book)

	err := GetBookById(book, id)
	
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(models.Error{
			Message: "No se encontró el libro",
			Code:    fiber.StatusNotFound,
		})
	}

	if result := database.DB.Preload("Reviews").Delete(book); result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Error{
			Message: "No se pudo eliminar el libro",
			Code:    fiber.StatusInternalServerError,
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}