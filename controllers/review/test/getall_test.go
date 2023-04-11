package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Daizaikun/back-library/controllers/review/handlers"
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetAllReview(t *testing.T) {

	// Configurar la base de datos
	database.DB = database.Connect()
	database.DB.AutoMigrate(&models.Review{})

	user := models.User{
		Email:    "jasom@asfsdasdf",
		Password: "asdasdasd",
		Name:     "asdasdasd",
	}
	database.DB.Create(&user)

	book := models.Book{
		Title:  "Pruebaaaaaa",
		Resume: "Pruebaaaaaa",
	}
	database.DB.Create(&book)

	// Crear algunos reviews de prueba
	reviews := []models.Review{
		{
			Comment: "Autor 1",
			Score:  4,
			UserID: user.ID,
			BookID: book.ID,
		},
		{
			Comment: "Autor 4",
			Score: 4,
			UserID: user.ID,
			BookID: book.ID},
		{
			Comment: "Autor 3",
			Score:  4,
			UserID: user.ID,
			BookID: book.ID,
		},
	}

	// Limpiar la base de datos
	defer database.DB.Where(user).Delete(&user)
	defer database.DB.Where(book).Delete(&book)
	defer delete(reviews)

	// Crear una solicitud HTTP para obtener todos los reviews
	request := httptest.NewRequest(http.MethodGet, "/crud/author", nil)

	// Ejecutar la solicitud HTTP utilizando Fiber
	app := fiber.New()
	app.Get("/crud/author", handlers.GetAll)

	responseOrigin, err := app.Test(request, -1)
	if err != nil {
		t.Error(err)
	}

	for _, review := range reviews {
		result := database.DB.Create(&review)
		if result.Error != nil {
			t.Error(result.Error)
		}
	}

	response, err := app.Test(request, -1)
	if err != nil {
		t.Error(err)
	}

	// Leer el cuerpo de la respuesta HTTP original
	var responseReviewOrigin []models.Review
	json.NewDecoder(responseOrigin.Body).Decode(&responseReviewOrigin)

	// Verificar que la respuesta HTTP sea correcta
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Leer el cuerpo de la respuesta HTTP
	var responseReview []models.Review
	json.NewDecoder(response.Body).Decode(&responseReview)

	// Verificar que se devuelvan todos los reviews
	assert.Equal(t, len(reviews), len(responseReview)-len(responseReviewOrigin))

}

func delete(reviews []models.Review) {
	for _, review := range reviews {
		database.DB.Where(review).Delete(&review)
	}
}
