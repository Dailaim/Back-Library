package tests

import (

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

func TestDeleteAuthor(t *testing.T) {
	// Configurar la base de datos
	database.DB = database.Connect()
	database.DB.AutoMigrate(&models.Author{})

	// Crear un autor de prueba
	author := models.Author{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}
	database.DB.Create(&author)

	// Crear una solicitud HTTP para eliminar el autor
	request := httptest.NewRequest(http.MethodDelete, "/crud/author/"+fmt.Sprint(author.ID), nil)

	// Ejecutar la solicitud HTTP utilizando Fiber
	app := fiber.New()
	app.Delete("/crud/author/:id", handlers.Delete)
	response, err := app.Test(request, -1)
	if err != nil {
		t.Error(err)
	}

	// Verificar que la respuesta HTTP sea correcta
	assert.Equal(t, http.StatusNoContent, response.StatusCode)

	// Verificar que el autor haya sido eliminado de la base de datos
	var count int64
	database.DB.Model(&models.Author{}).Where("id = ?", author.ID).Count(&count)
	assert.Equal(t, int64(0), count)
}