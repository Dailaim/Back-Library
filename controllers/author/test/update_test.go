package tests

import (
	"bytes"
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

func TestUpdateAuthor(t *testing.T) {
	// Configurar la base de datos
	database.DB = database.Connect()
	database.DB.AutoMigrate(&models.Author{})

	// Crear un autor de prueba
	author := models.Author{
		Name:  "Pruebaaaaaa",
	}
	database.DB.Create(&author)
	defer database.DB.Where(author).Delete(&author)

	// Actualizar el autor de prueba
	author.Name = "Prueba actualizada"

	// Crear una solicitud HTTP con el autor de prueba actualizado
	requestBody, err := json.Marshal(author)
	if err != nil {
		t.Error(err)
	}

	request := httptest.NewRequest(http.MethodPut, "/crud/author/"+ fmt.Sprint(author.ID), bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	// Ejecutar la solicitud HTTP utilizando Fiber
	app := fiber.New()
	app.Put("/crud/author/:id", handlers.Update)
	response, err := app.Test(request, -1)
	if err != nil {
		t.Error(err)
	}

	// Verificar que la respuesta HTTP sea correcta
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Leer el cuerpo de la respuesta HTTP
	var responseAuthor models.Author
	json.NewDecoder(response.Body).Decode(&responseAuthor)

	// Verificar que el nombre del autor actualizado sea correcto
	assert.Equal(t, author.Name, responseAuthor.Name)
}