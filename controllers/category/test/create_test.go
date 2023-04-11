package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Daizaikun/back-library/controllers/category/handlers"
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateCategory(t *testing.T) {
	// Configurar la base de datos
	database.DB = database.Connect()
	database.DB.AutoMigrate(&models.Category{})

	// Crear un autor de prueba
	category := models.Category{
		Name:  "Pruebaaaaaa",
	}
	defer database.DB.Where(category).Delete(&category)

	// Crear una solicitud HTTP con el category de prueba
	requestBody, err := json.Marshal(category)
	if err != nil {
		t.Error(err)
	}

	request := httptest.NewRequest(http.MethodPost, "/crud/category", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	// Ejecutar la solicitud HTTP utilizando Fiber
	app := fiber.New()
	app.Post("/crud/category", handlers.Create)
	response, err := app.Test(request, -1)
	if err != nil {
		t.Error(err)
	}

	// Verificar que la respuesta HTTP sea correcta
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Leer el cuerpo de la respuesta HTTP
	var responseCategory models.Category
	json.NewDecoder(response.Body).Decode(&responseCategory)

	// Verificar que el autor tenga un ID
	assert.NotEqual(t, uint(0), responseCategory.ID)
}