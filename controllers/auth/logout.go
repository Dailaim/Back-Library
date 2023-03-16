package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) error {
    // Borrar la cookie del token de autenticación
    c.Cookie(&fiber.Cookie{
        Name:     "jwt",
        Value:    "",
        Expires:  time.Now().Add(-time.Hour),
        HTTPOnly: true,
        Secure:   true,
    })

    // Devolver una respuesta JSON con un mensaje de éxito
    return c.JSON(fiber.Map{
        "message": "Successfully logged out",
    })
}