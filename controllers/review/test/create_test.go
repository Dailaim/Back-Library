package tests

import (
	"bytes"
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

func TestCreateReview(t *testing.T) {
	// Configurar la base de datos
	database.DB = database.Connect()
	database.DB.AutoMigrate(&models.Review{})

	user := models.User{
		Email: "jasom@asfsdasdf",
		Password: "asdasdasd",
		Name: "asdasdasd",

	}
	database.DB.Create(&user)

	book := models.Book{
		Title:  "Pruebaaaaaa",
		Resume: "Pruebaaaaaa",
	}
	database.DB.Create(&book)

	// Crear un review de prueba
	review := models.Review{
		Comment: "Pruebaaaaaa",
		Score:   4,
		UserID:  user.ID,
		BookID:  book.ID,
	}

	defer database.DB.Where(user).Delete(&user)
	defer database.DB.Where(book).Delete(&book)
	defer database.DB.Where(review).Delete(&review)

	// Crear una solicitud HTTP con el review de prueba
	requestBody, err := json.Marshal(review)
	if err != nil {
		t.Error(err)
	}

	request := httptest.NewRequest(http.MethodPost, "/crud/review", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	// Ejecutar la solicitud HTTP utilizando Fiber
	app := fiber.New()
	app.Post("/crud/review", handlers.Create)
	response, err := app.Test(request, -1)
	if err != nil {
		t.Error(err)
	}

	// Verificar que la respuesta HTTP sea correcta
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Leer el cuerpo de la respuesta HTTP
	var responseReview models.Review
	json.NewDecoder(response.Body).Decode(&responseReview)

	// Verificar que el review tenga un ID
	assert.NotEqual(t, uint(0), responseReview.ID)
}
