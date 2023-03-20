package app

import (
	
	"github.com/Daizaikun/back-library/router"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func App() *fiber.App {

	//crear la instancia de la app

	app := fiber.New()

	// middleware

	app.Use(cors.New())

	app.Use(recover.New())

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// Crear routs de la aplicación

	router.CRUD(app.Group("/crud"))

	//sirve las images de la aplicación

	router.Image(app.Group("/images/"))

	router.Auth(app.Group("/auth"))

	return app
}
