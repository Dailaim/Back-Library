package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"github.com/Daizaikun/back-library/app/middleware"
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
)

func HandleRegistration(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	// Validar que el email y la contraseña no estén vacíos
	if user.Email == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "email and password are required",
		})
	}

	// Validar que el email no esté en uso
	existingUser := models.User{}
	result := database.DB.Where("email = ?", user.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "email already in use",
		})
	}

	// Hashear la contraseña del usuario
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Guardar el usuario en la base de datos
	result = database.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	// Establecer el token de acceso en la estructura User
	user.AccessToken, err = generateToken(user.ID)

	if err != nil {
		return err
	}
	user.Password = ""

	// Devolver la estructura User al cliente
	return c.JSON(user)
}

// Generar el token de acceso
func generateToken(id uint) (string, error) {

	claims := jwt.MapClaims{}
	claims["user_id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(middleware.SecretKey))

}
