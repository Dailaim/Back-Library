package auth

import (
	userCRUD "github.com/Daizaikun/back-library/controllers/user"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func NewPassword(c *fiber.Ctx) error {
	type Input struct {
		Email       string `json:"email"`
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	var input Input
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	// Obtener al usuario de la base de datos
	user := new(models.User)
	if err := userCRUD.GetUserByEmail(user, input.Email); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	// Verificar la contraseña antigua
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.OldPassword)); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	// Encriptar la nueva contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to generate hashed password")
	}

	// Actualizar la contraseña en la base de datos
	user.Password = string(hashedPassword)
	if err := userCRUD.UpdateUser(user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update password")
	}

	return c.SendStatus(fiber.StatusOK)
} 