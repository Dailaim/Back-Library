package handlers

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

// @Summary Update author
// @Tags Author
// @Description Update author
// @Produce json
// @Param id path string true "Author ID"
// @Success 200 {object} models.Author
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /crud/author/{id} [put]
func Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	author := new(models.Author)

	err := GetAuthorById(author, id)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(models.Error{
			Message: "No se encontró el autor",
			Code:    fiber.StatusNotFound,
			Error:   err,
		})
	}

	if err := ctx.BodyParser(author); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.Error{
			Message: "Error al actualizar el autor",
			Code:    fiber.StatusBadRequest,
			Error:   err,
		})
	}
	
	if err := UpdateAuthor(author); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Error{
			Message: "Error al actualizar el autor",
			Code:    fiber.StatusInternalServerError,
			Error:   err,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(author)
}


func UpdateAuthor(author *models.Author) error {
	result := database.DB.Save(author)
	return result.Error
}
