package handlers

import (
	AuthorModels "github.com/Daizaikun/back-library/controllers/author/models"
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

// @Summary Update author
// @Tags Author
// @Description Update author
// @Produce json
// @Param id path string true "Author ID"
// @Success 200 {object} AuthorModels.SingleAuthorResponse
// @Failure 400 {object} AuthorModels.Response
// @Failure 404 {object} AuthorModels.Response
// @Failure 500 {object} AuthorModels.Response
// @Router /crud/author/{id} [put]
func Update(ctx *fiber.Ctx) error {
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

	updateAuthor := new(AuthorModels.NewAuthor)

	if err := ctx.BodyParser(updateAuthor); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(AuthorModels.Response{
			Error: &AuthorModels.Error{
				Message: "Error al parsear el body",
				Code:    fiber.StatusBadRequest,
			},
			Data: nil,
		})
	}

	author = &models.Author{
		FirstName: updateAuthor.FirstName,
		LastName: updateAuthor.LastName,
		Age: updateAuthor.Age,
	}
	
	if err := UpdateAuthor(author); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Error{
			Message: "Error al actualizar el autor",
			Code:    fiber.StatusInternalServerError,
			Error:   err,
		})
	}

	AuthorResponse := AuthorModels.Author{
		ID:        author.ID,
		FirstName: author.FirstName,
		LastName:  author.LastName,
		Age:       author.Age,
	}



	return ctx.Status(fiber.StatusOK).JSON(AuthorModels.Response{
		Error: nil,
		Data:  AuthorResponse,
	})
}


func UpdateAuthor(author *models.Author) error {
	result := database.DB.Save(author)
	return result.Error
}
