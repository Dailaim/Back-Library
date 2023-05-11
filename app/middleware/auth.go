package middleware

import (
	"fmt"
	"log"
	"strings"

	"github.com/Daizaikun/back-library/helpers"
	"github.com/gofiber/fiber/v2"
)

const SecretKey = "library"

func AuthMiddleware(c *fiber.Ctx) error {
	// Obtener el token de acceso del encabezado Authorization
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "no authorization token provided",
		})
	}

	// Check that the Authorization header is in the correct format.
	if !strings.HasPrefix(authHeader, "Bearer") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization header must be in the format of Bearer [token]",
		})
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Check if the token is blacklisted.
	isBlackListed := IsTokenBlacklisted(tokenString)
	if isBlackListed {
		log.Printf("Token is blacklisted: %s", tokenString)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid authorization token",
		})
	}

	claims, err := helpers.ValidateToken(tokenString)

	if err != nil {
		log.Printf("JWT Token validation. %v", err)

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid authorization token",
		})
	}

	userID := fmt.Sprintf("%v", claims["user_id"])
	c.Set("user_id", userID)

	return c.Next()
}
