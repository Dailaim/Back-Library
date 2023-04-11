package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Daizaikun/back-library/controllers/book/handlers"
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestUpdateBook(t *testing.T) {
	// Configurar la base de datos
	database.DB = database.Connect()
	database.DB.AutoMigrate(&models.Book{})

	// Crear un autor de prueba
	book := models.Book{
		Title: "Pruebaaaaaa",
	}
	database.DB.Create(&book)
	defer database.DB.Where(book).Delete(&book)

	// Actualizar el autor de prueba
	book.Title = "Prueba actualizada"

	// Crear una solicitud HTTP con el autor de prueba actualizado
	requestBody, err := json.Marshal(book)
	if err != nil {
		t.Error(err)
	}

	request := httptest.NewRequest(http.MethodPut, "/crud/book/"+fmt.Sprint(book.ID), bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	// Ejecutar la solicitud HTTP utilizando Fiber
	app := fiber.New()
	app.Put("/crud/book/:id", handlers.Update)
	response, err := app.Test(request, -1)
	if err != nil {
		t.Error(err)
	}

	// Verificar que la respuesta HTTP sea correcta
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Leer el cuerpo de la respuesta HTTP
	var responseBook models.Book
	json.NewDecoder(response.Body).Decode(&responseBook)

	// Verificar que el nombre del autor actualizado sea correcto
	assert.Equal(t, book.Title, responseBook.Title)
}
