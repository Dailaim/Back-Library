package auth

import (
	"time"

	userCRUD "github.com/Daizaikun/back-library/controllers/user"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)


func Login(c *fiber.Ctx) error {
    type LoginInput struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    var input LoginInput
    if err := c.BodyParser(&input); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
    }

	user := new(models.User)

    // Buscar al usuario en la base de datos
    err := userCRUD.GetUserByEmail(user, input.Email)
    if err != nil {
        return fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
    }

    // Verificar la contraseña
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

    // Crear el token JWT
    claims := &Claims{
        Email: user.Email,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // El token expira en 24 horas
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte("library"))
    if err != nil {
        return fiber.NewError(fiber.StatusInternalServerError, "Failed to create token")
    }

    return c.JSON(fiber.Map{
        "token": tokenString,
    })
}