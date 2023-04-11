package tests

import (

	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Daizaikun/back-library/controllers/review/handlers"
	"github.com/Daizaikun/back-library/database"       
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestDeleteReview(t *testing.T) {
	// Configurar la base de datos
	database.DB = database.Connect()
	database.DB.AutoMigrate(&models.Review{})

	// Crear un review de prueba
	user := models.User{
		Email: "jasom@asfsdasdf",
		Password: "asdasdasd",
		Name: "asdasdasd",

	}
	database.DB.Create(&user)

	book := models.Book{
		Title:  "Pruebaaaaaa",
		Resume: "Pruebaaaaaa",
	}
	database.DB.Create(&book)

	// Crear un review de prueba
	review := models.Review{
		Comment: "Pruebaaaaaa",
		Score:   4,
		UserID:  user.ID,
		BookID:  book.ID,
	}
	database.DB.Create(&review)

	defer database.DB.Where(user).Delete(&user)
	defer database.DB.Where(book).Delete(&book)
	defer database.DB.Where(review).Delete(&review)

	// Crear una solicitud HTTP para eliminar el review
	request := httptest.NewRequest(http.MethodDelete, "/crud/review/"+fmt.Sprint(review.ID), nil)

	// Ejecutar la solicitud HTTP utilizando Fiber
	app := fiber.New()
	app.Delete("/crud/review/:id", handlers.Delete)
	response, err := app.Test(request, -1)
	if err != nil {
		t.Error(err)
	}

	// Verificar que la respuesta HTTP sea correcta
	assert.Equal(t, http.StatusNoContent, response.StatusCode)

	// Verificar que el review haya sido eliminado de la base de datos
	var count int64
	database.DB.Where("id = ?", review.ID).Count(&count)
	assert.Equal(t, int64(0), count)
}