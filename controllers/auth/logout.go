package auth

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"

	"github.com/Daizaikun/back-library/app/middleware"
)

func HandleLogout(c *fiber.Ctx) error {
	// Obtener el token de acceso del encabezado Authorization
	authHeader := c.Get("Authorization")
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	// Invalidar el token de acceso
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(middleware.SecretKey), nil
	})
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

	// Agregar el token a la lista negra para invalidarlo
	err = middleware.AddToken(tokenString)
	if err != nil {
		return err
	}

	// Devolver una respuesta vacía con un estado HTTP 204 No Content
	return c.SendStatus(fiber.StatusNoContent)
}
