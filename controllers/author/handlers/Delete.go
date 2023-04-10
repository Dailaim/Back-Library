package handlers

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

// @Summary Delete author
// @Tags Author
// @Description Delete author
// @Produce json
// @Param id path string true "Author ID"
// @Success 204 
// @Failure 400
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /crud/author/{id} [delete]
func Delete(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	author := new(models.Author)

	err := GetAuthorById(author, id)
	
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(models.Error{
			Message: "No se encontró el autor",
			Code:    fiber.StatusNotFound,
			Error:  err,
		})
	}

	if result := database.DB.Delete(author); result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Error{
			Message: "No se pudo eliminar el autor",
			Code:    fiber.StatusInternalServerError,
			Error:  result.Error,
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}