package middleware

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)


func AuthMiddleware(c *fiber.Ctx) error {
    // Obtener el token de acceso del encabezado Authorization
    authHeader := c.Get("Authorization")
    if authHeader == "" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "no authorization token provided",
        })
    }

    // Extraer el token de acceso del encabezado Authorization
    tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

    // Validar el token de acceso
    claims := jwt.MapClaims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte("secret_key"), nil
    })

     // Verificar si el token de acceso está en la lista negra

    if IsTokenBlacklisted(tokenString) {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "invalid authorization token",
        })
    }



    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "invalid authorization token",
        })
    }
    if !token.Valid {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "invalid authorization token",
        })
    }

    // Obtener el ID de usuario del token de acceso
    userID := claims["user_id"].(string)

    // Establecer el ID de usuario en el contexto de Fiber
    c.Set("user_id", userID)

    // Llamar a la función siguiente en la cadena de middleware
    return c.Next()
}