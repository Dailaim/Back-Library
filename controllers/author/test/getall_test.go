package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Daizaikun/back-library/controllers/author/handlers"
	AuthorModels "github.com/Daizaikun/back-library/controllers/author/models"
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetAllAuthors(t *testing.T) {
	
	// Configurar la base de datos
	database.DB = database.Connect()
	database.DB.AutoMigrate(&models.Author{})

	// Crear algunos autores de prueba
	authors := []models.Author{
		{FirstName: "Autor 1", LastName: "Apellido 1", Age: 30},
		{FirstName: "Autor 2", LastName: "Apellido 2", Age: 30},
		{FirstName: "Autor 3", LastName: "Apellido 3", Age: 30},
	}

	// Limpiar la base de datos
	defer delete(authors)

	// Crear una solicitud HTTP para obtener todos los autores
	request := httptest.NewRequest(http.MethodGet, "/crud/author", nil)

	// Ejecutar la solicitud HTTP utilizando Fiber
	app := fiber.New()
	app.Get("/crud/author", handlers.GetAll)

	responseOrigin, err := app.Test(request, -1)
	if err != nil {
		t.Error(err)
	}

	for _, author := range authors {
		result := database.DB.Create(&author)
		if result.Error != nil {
			t.Error(result.Error)
		}
	}

	response, err := app.Test(request, -1)
	if err != nil {
		t.Error(err)
	}

	// Leer el cuerpo de la respuesta HTTP original
	var responseAutorOrigin AuthorModels.MultipleAuthorsResponse
	json.NewDecoder(responseOrigin.Body).Decode(&responseAutorOrigin)

	// Verificar que la respuesta HTTP sea correcta
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Leer el cuerpo de la respuesta HTTP
	var responseAuthors AuthorModels.MultipleAuthorsResponse
	json.NewDecoder(response.Body).Decode(&responseAuthors)

	// Verificar que se devuelvan todos los autores
	assert.Equal(t, len(authors) , len(responseAuthors.Data) - len(responseAutorOrigin.Data))
}

func delete(authors []models.Author) {
	for _, author := range authors {
		database.DB.Where(author).Delete(&author)
	}
}
