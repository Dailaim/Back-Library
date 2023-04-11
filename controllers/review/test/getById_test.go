package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Daizaikun/back-library/controllers/review/handlers"
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetReviewById(t *testing.T) {
	// Configurar la base de datos
	database.DB = database.Connect()
	database.DB.AutoMigrate(&models.Review{})

	// Crear un review de prueba
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
	database.DB.Create(&review)

	defer database.DB.Where(user).Delete(&user)
	defer database.DB.Where(book).Delete(&book)
	defer database.DB.Where(review).Delete(&review)

	// Crear una solicitud HTTP para obtener el review
	request := httptest.NewRequest(http.MethodGet, "/crud/review/"+ fmt.Sprint(review.ID), nil)

	// Ejecutar la solicitud HTTP utilizando Fiber
	app := fiber.New()
	app.Get("/crud/review/:id", handlers.GetById)
	response, err := app.Test(request, -1)
	if err != nil {
		t.Error(err)
	}

	// Verificar que la respuesta HTTP sea correcta
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Leer el cuerpo de la respuesta HTTP
	var responseReview models.Review
	json.NewDecoder(response.Body).Decode(&responseReview)

	// Verificar que el review tenga el mismo ID que el review de prueba
	assert.Equal(t, review.ID, responseReview.ID)
}