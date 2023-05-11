package handlers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	AuthModels "github.com/Daizaikun/back-library/controllers/auth/models"
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/helpers"
	"github.com/Daizaikun/back-library/models"
)

// @Summary Login
// @Tags Auth
// @Description Login
// @accept json
// @Produce json
// @Param user body AuthModels.UserLogin true "User"
// @Success 200 {object} AuthModels.Response "User"
// @Failure 401 {object} AuthModels.Response
// @Failure 500 {object} AuthModels.Response
// @Router /auth/login [post]
func Authentication(c *fiber.Ctx) error {

	var user AuthModels.UserLogin
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	var existingUser models.User
	// Buscar el usuario en la base de datos
	result := database.DB.Where("email = ?", user.Email).First(&existingUser)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(AuthModels.Response{
			Data: nil,
			Error: &AuthModels.Error{
				Message: "El email o la contraseña son incorrectos",
				Code:    fiber.StatusUnauthorized,
			},
		},
		)
	}

	// Validar la contraseña del usuario
	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(AuthModels.Response{
			Data: nil,
			Error: &AuthModels.Error{
				Message: "El email o la contraseña son incorrectos",
				Code:    fiber.StatusUnauthorized,
			},
		},
		)
	}

	// Establecer el token de acceso en la estructura User
	tokenAccess, err := helpers.GenerateToken(existingUser.ID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(AuthModels.Response{
			Data: nil,
			Error: &AuthModels.Error{
				Message: "Error al generar el token",
				Code:    fiber.StatusInternalServerError,
			},
		})
	}

	response := AuthModels.Response{
		Data: &AuthModels.Data{
			TokenAccess: tokenAccess,
		},
		Error: nil,
	}

	// Devolver la estructura User al cliente
	return c.Status(fiber.StatusOK).JSON(response)
}
