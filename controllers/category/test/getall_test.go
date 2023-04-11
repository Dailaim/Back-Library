package tests

import (
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

func TestGetAllCategories(t *testing.T) {
	
	// Configurar la base de datos
	database.DB = database.Connect()
	database.DB.AutoMigrate(&models.Category{})

	// Crear algunos autores de prueba
	categories := []models.Category{
		{Name: "Autor 1"},
		{Name: "Autor 4"},
		{Name: "Autor 3"},
	}

	// Limpiar la base de datos
	defer delete(categories)

	// Crear una solicitud HTTP para obtener todos los autores
	request := httptest.NewRequest(http.MethodGet, "/crud/category", nil)

	// Ejecutar la solicitud HTTP utilizando Fiber
	app := fiber.New()
	app.Get("/crud/category", handlers.GetAll)

	responseOrigin, err := app.Test(request, -1)
	if err != nil {
		t.Error(err)
	}

	for _, category := range categories {
		result := database.DB.Create(&category)
		if result.Error != nil {
			t.Error(result.Error)
		}
	}

	response, err := app.Test(request, -1)
	if err != nil {
		t.Error(err)
	}

	// Leer el cuerpo de la respuesta HTTP original
	var responseCategoriesOrigin []models.Author
	json.NewDecoder(responseOrigin.Body).Decode(&responseCategoriesOrigin)

	// Verificar que la respuesta HTTP sea correcta
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Leer el cuerpo de la respuesta HTTP
	var responseCategories []models.Author
	json.NewDecoder(response.Body).Decode(&responseCategories)

	// Verificar que se devuelvan todos los autores
	assert.Equal(t, len(categories) , len(responseCategories) - len(responseCategoriesOrigin))

}

func delete(categories []models.Category) {
	for _, category := range categories {
		database.DB.Where(category).Delete(&category)
	}
}
