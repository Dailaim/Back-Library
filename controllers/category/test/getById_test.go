package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Daizaikun/back-library/controllers/category/handlers"
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetCategoryById(t *testing.T) {
	// Configurar la base de datos
	database.DB = database.Connect()
	database.DB.AutoMigrate(&models.Category{})

	// Crear un category de prueba
	category := models.Category{
		Name:  "Pruebaaaaaa",
	}
	database.DB.Create(&category)
	defer database.DB.Where(category).Delete(&category)

	// Crear una solicitud HTTP para obtener el category
	request := httptest.NewRequest(http.MethodGet, "/crud/category/"+ fmt.Sprint(category.ID), nil)

	// Ejecutar la solicitud HTTP utilizando Fiber
	app := fiber.New()
	app.Get("/crud/category/:id", handlers.GetById)
	response, err := app.Test(request, -1)
	if err != nil {
		t.Error(err)
	}

	// Verificar que la respuesta HTTP sea correcta
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Leer el cuerpo de la respuesta HTTP
	var responseCategory models.Category
	json.NewDecoder(response.Body).Decode(&responseCategory)

	// Verificar que el autor tenga el mismo ID que el autor de prueba
	assert.Equal(t, category.ID, responseCategory.ID)
}