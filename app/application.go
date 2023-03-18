package application

import (
	"fmt"
	"strings"

	"github.com/Daizaikun/back-library/router"
	"github.com/gofiber/fiber/v2"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func App() *fiber.App {

	//crear la instancia de la app

	app := fiber.New()

	//descifra los datos

	// middleware

	app.Use(cors.New())

	app.Use(recover.New())

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// Crear rout's de la aplicación

	router.CRUD(app.Group("/crud"))

	//sirve las images de la aplicación

	router.Image(app.Group("/images/"))

	app.Use(func(c *fiber.Ctx) error {
		// Obtenemos el token JWT de los headers
		authHeader := c.Get("Authorization")
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// Validamos el token JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verificamos si el algoritmo de cifrado es HMAC y si la clave es válida
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("library"), nil
		})
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		// Si el token es válido, seteamos el usuario en el contexto
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}
		userID := claims["user_id"].(string)
		c.Set("user_id", userID)

		return c.Next()
	})

	router.Auth(app.Group("/auth"))

	return app
}
