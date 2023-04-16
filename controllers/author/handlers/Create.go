package handlers

import (
	"github.com/Daizaikun/back-library/database"
	DatabaseModels "github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"

	AuthorModels "github.com/Daizaikun/back-library/controllers/author/models"
)

// @Summary Create author
// @Tags Author
// @Description Create author
// @Produce json
// @Param Author body AuthorModels.NewAuthor true "Author data"
// @Success 200 {object} AuthorModels.SingleAuthorResponse
// @Failure 400 {object} AuthorModels.Response
// @Failure 500 {object} AuthorModels.Response
// @Router /crud/author [post]
func Create(ctx *fiber.Ctx) error {
	// Crear un nuevo usuario
	Author := new(AuthorModels.NewAuthor)

	// Parsear el body de la petición
	err := ctx.BodyParser(Author)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(DatabaseModels.Error{
			Message: "No se pudo crear el autor",
			Code:    fiber.StatusBadRequest,
			Error:   err,
		})
	}

	if Author.FirstName == "" || Author.LastName == "" || Author.Age == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(AuthorModels.Response{
			Error: &AuthorModels.Error{
				Message: "No se pudo crear el autor con campos vacíos",
				Code:    fiber.StatusBadRequest,
			},
			Data: nil,
		})
	}

	NewAuthor := new(DatabaseModels.Author)

	// Crear el objeto de la base de datos
	NewAuthor = &DatabaseModels.Author{
		FirstName: Author.FirstName,
		LastName:  Author.LastName,
		Age:       Author.Age,
	}

	// Comprobar si el autor ya existe
	result := database.DB.Where("name = ?", NewAuthor)
	if result.RowsAffected > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(AuthorModels.Response{
			Error: &AuthorModels.Error{
				Message: "El autor ya existe",
				Code:    fiber.StatusBadRequest,
			},
			Data: nil,
		})
	}

	// Crear el review en la base de datos
	if result := database.DB.Create(&NewAuthor); result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(AuthorModels.Response{
			Error: &AuthorModels.Error{
				Message: "No se pudo crear el autor",
				Code:    fiber.StatusInternalServerError,
			},
			Data: nil,
		})
	}

	// Devolver la respuesta JSON con el objeto creado y un estado HTTP 200
	return ctx.Status(fiber.StatusOK).JSON(AuthorModels.Response{
		Error: nil,
		Data:  NewAuthor,
	})
}
