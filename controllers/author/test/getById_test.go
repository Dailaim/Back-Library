package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Daizaikun/back-library/controllers/author/handlers"
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetAuthorById(t *testing.T) {
	// Configurar la base de datos
	database.DB = database.Connect()
	database.DB.AutoMigrate(&models.Author{})

	// Crear un autor de prueba
	author := models.Author{
		Name:  "Pruebaaaaaa",
	}
	database.DB.Create(&author)
	defer database.DB.Where(author).Delete(&author)

	// Crear una solicitud HTTP para obtener el autor
	request := httptest.NewRequest(http.MethodGet, "/crud/author/"+ fmt.Sprint(author.ID), nil)

	// Ejecutar la solicitud HTTP utilizando Fiber
	app := fiber.New()
	app.Get("/crud/author/:id", handlers.GetById)
	response, err := app.Test(request, -1)
	if err != nil {
		t.Error(err)
	}

	// Verificar que la respuesta HTTP sea correcta
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Leer el cuerpo de la respuesta HTTP
	var responseAuthor models.Author
	json.NewDecoder(response.Body).Decode(&responseAuthor)

	// Verificar que el autor tenga el mismo ID que el autor de prueba
	assert.Equal(t, author.ID, responseAuthor.ID)
}