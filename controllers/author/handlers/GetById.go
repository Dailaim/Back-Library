package handlers

import (
	AuthorModels "github.com/Daizaikun/back-library/controllers/author/models"
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

// @Summary Get author by id
// @Tags Author
// @Description Get author by id
// @Produce json
// @Param id path string true "Author ID"
// @Success 200 {object} AuthorModels.SingleAuthorResponse
// @Failure 400
// @Failure 404 {object} AuthorModels.Response
// @Failure 500 {object} AuthorModels.Response
// @Router /crud/author/{id} [get]
func GetById(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	author := new(models.Author)

	err := GetAuthorById(author, id)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(AuthorModels.Response{
			Error: &AuthorModels.Error{
				Message: "No se encontró el autor",
				Code:    fiber.StatusNotFound,
			},
			Data: nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(AuthorModels.Response{
		Error: nil,
		Data:  author,
	})
}

func GetAuthorById(author *models.Author, id string) error {

	result := database.DB.First(author, id)

	return result.Error
}
