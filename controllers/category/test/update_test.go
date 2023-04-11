package tests

import (
	"bytes"
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

func TestUpdateCategory(t *testing.T) {
	
	// Configurar la base de datos
	database.DB = database.Connect()
	database.DB.AutoMigrate(&models.Category{})

	// Crear un autor de prueba
	category := models.Category{
		Name:  "Pruebaaaaaa",
	}
	database.DB.Create(&category)
	defer database.DB.Where(category).Delete(&category)

	// Actualizar el autor de prueba
	category.Name = "Prueba actualizada"

	// Crear una solicitud HTTP con el autor de prueba actualizado
	requestBody, err := json.Marshal(category)
	if err != nil {
		t.Error(err)
	}

	request := httptest.NewRequest(http.MethodPut, "/crud/category/"+ fmt.Sprint(category.ID), bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	// Ejecutar la solicitud HTTP utilizando Fiber
	app := fiber.New()
	app.Put("/crud/category/:id", handlers.Update)
	response, err := app.Test(request, -1)
	if err != nil {
		t.Error(err)
	}

	// Verificar que la respuesta HTTP sea correcta
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Leer el cuerpo de la respuesta HTTP
	var responseCategory models.Category
	json.NewDecoder(response.Body).Decode(&responseCategory)

	// Verificar que el nombre del autor actualizado sea correcto
	assert.Equal(t, category.Name, responseCategory.Name)
}