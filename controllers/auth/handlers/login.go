package handlers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"

)

// @Summary Login
// @Tags Auth
// @Description Login
// @accept json
// @Produce json
// @Param user body models.User true "User"
// @Success 200 {object} models.User "User" 
// @Failure 401 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /auth/login [post]
func Authentication(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	var existingUser models.User

	// Buscar el usuario en la base de datos
	result := database.DB.Where("email = ?", user.Email).First(&existingUser)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "incorrect email or password",
		})
	}

	// Validar la contraseña del usuario
	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "incorrect email or password",
		})
	}

	existingUser.Password = ""

	// Establecer el token de acceso en la estructura User
	existingUser.AccessToken, err = generateToken(existingUser.ID)

	if err != nil {
		return err
	}

	// Devolver la estructura User al cliente
	return c.JSON(existingUser)
}
