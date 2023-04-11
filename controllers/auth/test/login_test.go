package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"	
	"golang.org/x/crypto/bcrypt"

	"github.com/Daizaikun/back-library/controllers/auth/handlers"
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
)


func TestAuthentication(t *testing.T) {
	// Configurar la base de datos
	database.DB = database.Connect()
	database.DB.AutoMigrate(&models.User{})

	// Crear un usuario de prueba
	user := models.User{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "secret",
	}
	defer database.DB.Where("email = ?", user.Email).Delete(&user)

	// Hashear la contraseña del usuario de prueba
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		t.Error(err)
	}
	user.Password = string(hashedPassword)

	result := database.DB.Create(&user)
	if result.Error != nil {
		t.Error(result.Error)
	}

	// Crear una solicitud HTTP con el usuario de prueba
	requestBody := fmt.Sprintf(`{"email":"%s","password":"%s"}`, user.Email, "secret")
	request := httptest.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(requestBody))
	request.Header.Set("Content-Type", "application/json")

	// Ejecutar la solicitud HTTP utilizando Fiber
	app := fiber.New()
	app.Post("/auth/login", handlers.Authentication)
	response, err := app.Test(request, -1)
	if err != nil {
		t.Error(err)
	}

	// Verificar que la respuesta HTTP sea correcta
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Leer el cuerpo de la respuesta HTTP
	var responseUser models.User
	json.NewDecoder(response.Body).Decode(&responseUser)

	// Verificar que el usuario tenga un token de acceso
	assert.NotEmpty(t, responseUser.AccessToken)
}

