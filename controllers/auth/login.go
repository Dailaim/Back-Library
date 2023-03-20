package auth

import (

	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)


func HandleAuthentication(c *fiber.Ctx) error {
    var user models.User
    if err := c.BodyParser(&user); err != nil {
        return err
    }

    // Buscar el usuario en la base de datos
    result := database.DB.Where("email = ?", user.Email).First(&user)
    if result.RowsAffected == 0 {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "incorrect email or password",
        })
    }

    // Validar la contraseña del usuario
    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password))
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "incorrect email or password",
        })
    }

    // Generar el token de acceso
    claims := jwt.MapClaims{}
    claims["user_id"] = user.ID
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedToken, err := token.SignedString([]byte("secret_key"))

    if err != nil {
        return err
    }

    // Establecer el token de acceso en la estructura User
    user.AccessToken = signedToken

    // Devolver la estructura User al cliente
    return c.JSON(user)
}