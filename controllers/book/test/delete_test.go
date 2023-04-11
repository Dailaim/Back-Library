package tests

import (

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

func TestDeleteBook(t *testing.T) {
	// Configurar la base de datos
	database.DB = database.Connect()
	database.DB.AutoMigrate(&models.Book{})

	// Crear un autor de prueba
	Book := models.Book{
		Title: "Prueba",
		Resume: "Prueba",
	}
	database.DB.Create(&Book)

	// Crear una solicitud HTTP para eliminar el autor
	request := httptest.NewRequest(http.MethodDelete, "/crud/book/"+fmt.Sprint(Book.ID), nil)

	// Ejecutar la solicitud HTTP utilizando Fiber
	app := fiber.New()
	app.Delete("/crud/book/:id", handlers.Delete)
	response, err := app.Test(request, -1)
	if err != nil {
		t.Error(err)
	}

	// Verificar que la respuesta HTTP sea correcta
	assert.Equal(t, http.StatusNoContent, response.StatusCode)

	// Verificar que el autor haya sido eliminado de la base de datos
	var count int64
	database.DB.Model(&models.Book{}).Where("id = ?", Book.ID).Count(&count)
	assert.Equal(t, int64(0), count)
}