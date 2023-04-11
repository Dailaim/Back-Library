package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Daizaikun/back-library/controllers/book/handlers"
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateBook(t *testing.T) {

	// Configurar la base de datos
	database.DB = database.Connect()
	database.DB.AutoMigrate(&models.Book{})

	// Crear un book de prueba
	Book := models.Book{
	Title:  "Pruebaaaaaa",
	Resume: "Pruebaaaaaa",
	}
	defer database.DB.Where(Book).Delete(&Book)

	// Crear una solicitud HTTP con el autor de prueba
	requestBody, err := json.Marshal(Book)
	if err != nil {
		t.Error(err)
	}

	request := httptest.NewRequest(http.MethodPost, "/crud/book", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	// Ejecutar la solicitud HTTP utilizando Fiber
	app := fiber.New()
	app.Post("/crud/book", handlers.Create)
	response, err := app.Test(request, -1)
	if err != nil {
		t.Error(err)
	}

	// Verificar que la respuesta HTTP sea correcta
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Leer el cuerpo de la respuesta HTTP
	var responseBook models.Book
	json.NewDecoder(response.Body).Decode(&responseBook)

	// Verificar que el autor tenga un ID
	assert.NotEqual(t, uint(0), responseBook.ID)
}