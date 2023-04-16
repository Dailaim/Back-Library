package handlers

import (
	AuthorModels "github.com/Daizaikun/back-library/controllers/author/models"
	"github.com/Daizaikun/back-library/database"
	DatabaseModels "github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)

// @Summary Get all authors
// @Tags Author
// @Description Get all authors
// @Produce json
// @Success 200 {array} AuthorModels.MultipleAuthorsResponse
// @Failure 500 {object} AuthorModels.Response
// @Router /crud/author [get]
func GetAll(ctx *fiber.Ctx) error {
	authors := new([]DatabaseModels.Author)

	result := database.DB.Find(authors)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(AuthorModels.Response{
			Error: &AuthorModels.Error{
				Message: "No se pudieron obtener los autores",
				Code:    fiber.StatusInternalServerError,
			},
			Data: nil,
		})
	}
	var AuthorsResponse []AuthorModels.Author

	for _, author := range *authors {
		var Books []AuthorModels.Book
		for _, book := range author.Books {
			Books = append(Books, AuthorModels.Book{
				Title:  book.Title,
				Image:  book.Image,
				Resume: book.Resume,
			})
			

		}
		AuthorsResponse = append(AuthorsResponse, AuthorModels.Author{
			ID:        author.ID,
			FirstName: author.FirstName,
			LastName:  author.LastName,
			Age:       author.Age,
			Books:     Books,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(AuthorModels.MultipleAuthorsResponse{
		Error: nil,
		Data:  AuthorsResponse,
	})
}
