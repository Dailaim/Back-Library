package tests

import (

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

func TestDeleteCategory(t *testing.T) {

	// Configurar la base de datos
	database.DB = database.Connect()
	database.DB.AutoMigrate(&models.Category{})

	// Crear un category de prueba
	category := models.Category{
		Name: "Prueba",
	}
	database.DB.Create(&category)
	defer database.DB.Where(category).Delete(&category)

	// Crear una solicitud HTTP para eliminar el category
	request := httptest.NewRequest(http.MethodDelete, "/crud/category/"+fmt.Sprint(category.ID), nil)

	// Ejecutar la solicitud HTTP utilizando Fiber
	app := fiber.New()
	app.Delete("/crud/category/:id", handlers.Delete)
	response, err := app.Test(request, -1)
	if err != nil {
		t.Error(err)
	}

	// Verificar que la respuesta HTTP sea correcta
	assert.Equal(t, http.StatusNoContent, response.StatusCode)

	// Verificar que el autor haya sido eliminado de la base de datos
	var count int64
	database.DB.Model(&models.Category{}).Where("id = ?", category.ID).Count(&count)
	assert.Equal(t, int64(0), count)
}