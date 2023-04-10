package handlers

import (
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)


// @Summary Get author by id
// @Tags Author
// @Description Get author by id
// @Produce json
// @Param id path string true "Author ID"
// @Success 200 {object} models.Author
// @Failure 400
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /crud/author/{id} [get]
func GetById(ctx *fiber.Ctx) error {

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

	return ctx.Status(fiber.StatusOK).JSON(author)
}

func GetAuthorById(author *models.Author, id string) error {

	result := database.DB.First(author, id)

	return result.Error
}
