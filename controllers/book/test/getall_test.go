package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Daizaikun/back-library/controllers/book/handlers"
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetAllBook(t *testing.T) {
	
	// Configurar la base de datos
	database.DB = database.Connect()
	database.DB.AutoMigrate(&models.Book{})

	// Crear algunos autores de prueba
	Books := []models.Book{
		{Title: "Autor 1"},
		{Title: "Autor 4"},
		{Title: "Autor 3"},
	}

	// Limpiar la base de datos
	defer delete(Books)

	// Crear una solicitud HTTP para obtener todos los autores
	request := httptest.NewRequest(http.MethodGet, "/crud/book", nil)

	// Ejecutar la solicitud HTTP utilizando Fiber
	app := fiber.New()
	app.Get("/crud/book", handlers.GetAll)

	responseOrigin, err := app.Test(request, -1)
	if err != nil {
		t.Error(err)
	}

	for _, book := range Books {
		result := database.DB.Create(&book)
		if result.Error != nil {
			t.Error(result.Error)
		}
	}

	response, err := app.Test(request, -1)
	if err != nil {
		t.Error(err)
	}

	// Leer el cuerpo de la respuesta HTTP original
	var responseBookOrigin []models.Book
	json.NewDecoder(responseOrigin.Body).Decode(&responseBookOrigin)

	// Verificar que la respuesta HTTP sea correcta
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Leer el cuerpo de la respuesta HTTP
	var responseAuthors []models.Author
	json.NewDecoder(response.Body).Decode(&responseAuthors)

	// Verificar que se devuelvan todos los autores
	assert.Equal(t, len(Books) , len(responseAuthors) - len(responseBookOrigin))

}

func delete(books []models.Book) {
	for _, book := range books {
		database.DB.Where(book).Delete(&book)
	}
}
