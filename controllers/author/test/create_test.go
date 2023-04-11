package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Daizaikun/back-library/controllers/author/handlers"
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateAuthor(t *testing.T) {
	// Configurar la base de datos
	database.DB = database.Connect()
	database.DB.AutoMigrate(&models.Author{})

	// Crear un autor de prueba
	author := models.Author{
		Name:  "Pruebaaaaaa",
	}
	defer database.DB.Where(author).Delete(&author)

	// Crear una solicitud HTTP con el autor de prueba
	requestBody, err := json.Marshal(author)
	if err != nil {
		t.Error(err)
	}

	request := httptest.NewRequest(http.MethodPost, "/crud/author", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	// Ejecutar la solicitud HTTP utilizando Fiber
	app := fiber.New()
	app.Post("/crud/author", handlers.Create)
	response, err := app.Test(request, -1)
	if err != nil {
		t.Error(err)
	}

	// Verificar que la respuesta HTTP sea correcta
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Leer el cuerpo de la respuesta HTTP
	var responseAuthor models.Author
	json.NewDecoder(response.Body).Decode(&responseAuthor)

	// Verificar que el autor tenga un ID
	assert.NotEqual(t, uint(0), responseAuthor.ID)
}