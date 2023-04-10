package handlers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"github.com/Daizaikun/back-library/app/middleware"
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
)

// @Summary Register
// @Tags Auth
// @Description Register
// @Produce json
// @Param user body models.User true "User"
// @Success 200 {object} models.User
// @Failure 400	{object} models.Error
// @Failure 409 {object} models.Error
// @Router /auth/register [post]
func Registration(ctx *fiber.Ctx) error {
	var user models.User
	if err := ctx.BodyParser(&user); err != nil {
		return err
	}

	// Validar que el email y la contraseña no estén vacíos
	if user.Email == "" || user.Password == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.Error{
			Message: "El email y la contraseña son obligatorios",
			Code:   fiber.StatusBadRequest,
		})
	}

	// Validar que el email no esté en uso
	existingUser := models.User{}
	result := database.DB.Where("email = ?", user.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		return ctx.Status(fiber.StatusConflict).JSON(models.Error{
			Message: "El email ya está en uso",
			Code:   fiber.StatusConflict,
		})
	}

	// Hashear la contraseña del usuario
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Error{
			Message: "Error al hashear la contraseña",
			Code:   fiber.StatusInternalServerError,
			Error: err,
		})
	}
	user.Password = string(hashedPassword)

	// Guardar el usuario en la base de datos
	result = database.DB.Create(&user)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Error{
			Message: "Error al guardar el usuario",
			Code:   fiber.StatusInternalServerError,
			Error: result.Error,
		})
	}

	// Establecer el token de acceso en la estructura User
	user.AccessToken, err = generateToken(user.ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Error{
			Message: "Error al generar el token de acceso",
			Code:   fiber.StatusInternalServerError,
			Error: err,
		})
	}
	user.Password = ""

	// Devolver la estructura User al cliente
	return ctx.Status(fiber.StatusOK).JSON(user)
}

// Generar el token de acceso
func generateToken(id uint) (string, error) {

	claims := jwt.MapClaims{}
	claims["user_id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(middleware.SecretKey))

}
