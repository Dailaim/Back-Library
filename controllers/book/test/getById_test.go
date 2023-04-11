package tests

import (
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

func TestGetBookById(t *testing.T) {
	// Configurar la base de datos
	database.DB = database.Connect()
	database.DB.AutoMigrate(&models.Book{})

	// Crear un book de prueba
	book := models.Book{
		Title:  "Pruebaaaaaa",
	}
	database.DB.Create(&book)
	defer database.DB.Where(book).Delete(&book)

	// Crear una solicitud HTTP para obtener el Book
	request := httptest.NewRequest(http.MethodGet, "/crud/book/"+ fmt.Sprint(book.ID), nil)

	// Ejecutar la solicitud HTTP utilizando Fiber
	app := fiber.New()
	app.Get("/crud/book/:id", handlers.GetById)
	response, err := app.Test(request, -1)
	if err != nil {
		t.Error(err)
	}

	// Verificar que la respuesta HTTP sea correcta
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Leer el cuerpo de la respuesta HTTP
	var responseBook models.Book
	json.NewDecoder(response.Body).Decode(&responseBook)

	// Verificar que el autor tenga el mismo ID que el autor de prueba
	assert.Equal(t, book.ID, responseBook.ID)
}