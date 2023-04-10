package auth

import (
	"github.com/Daizaikun/back-library/controllers/auth/handlers"

	"github.com/gofiber/fiber/v2"
)

func Routers(app fiber.Router) {
	
	app.Post("/login", handlers.Authentication)
	app.Post("/register", handlers.Registration)
	app.Post("/logout", handlers.Logout)

}