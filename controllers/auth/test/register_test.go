package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Daizaikun/back-library/app/middleware"
	"github.com/Daizaikun/back-library/controllers/auth/handlers"
	AuthModels "github.com/Daizaikun/back-library/controllers/auth/models"
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRegistration(t *testing.T) {
	// Configurar la base de datos
	database.DB = database.Connect()
	database.DB.AutoMigrate(&models.User{})

	// Crear una solicitud HTTP con un usuario de prueba
	user := models.User{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "secret",
	}
	requestBody, err := json.Marshal(&user)
	if err != nil {
		t.Error(err)
	}
	defer database.DB.Where("email = ?", user.Email).Delete(&user)
	request := httptest.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	// Ejecutar la solicitud HTTP utilizando Fiber
	app := fiber.New()
	app.Post("/auth/register", handlers.Registration)
	response, err := app.Test(request, -1)
	if err != nil {
		t.Error(err)
	}

	// Verificar que la respuesta HTTP sea correcta
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Leer el cuerpo de la respuesta HTTP
	var responseUser AuthModels.Response
	err = json.NewDecoder(response.Body).Decode(&responseUser)
	if err != nil {
		t.Error(err)
	}

	// Verificar que el usuario tenga un token de acceso válido
	token, err := jwt.Parse(responseUser.Data.TokenAccess, func(token *jwt.Token) (interface{}, error) {
		return []byte(middleware.SecretKey), nil
	})
	if err != nil {
		t.Error(err)
	}

	assert.True(t, token.Valid)
}