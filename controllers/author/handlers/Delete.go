package handlers

import (
	AuthorModels "github.com/Daizaikun/back-library/controllers/author/models"
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

// @Summary Delete author
// @Tags Author
// @Description Delete author
// @Produce json
// @Param id path string true "Author ID"
// @Success 204 {object} AuthorModels.Response
// @Failure 400	
// @Failure 404 {object} AuthorModels.Response
// @Failure 500 {object} AuthorModels.Response
// @Router /crud/author/{id} [delete]
func Delete(ctx *fiber.Ctx) error {

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

	if result := database.DB.Delete(author); result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(AuthorModels.Response{
			Error: &AuthorModels.Error{
				Message: "No se pudo eliminar el autor",
				Code:    fiber.StatusInternalServerError,
			},
			Data: nil,
		})
	}

	return ctx.Status(fiber.StatusNoContent).JSON(AuthorModels.Response{
		Error: nil,
		Data:  nil,
	})	
}
