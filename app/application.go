package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	_ "github.com/Daizaikun/back-library/docs"
	"github.com/Daizaikun/back-library/router"

	"github.com/gofiber/swagger"
)

// @title Back Library API
// @version 1.0
// @description This is a sample server library server.
// @contact.name Daizaikun
// @contact.email laiglesias.min@gmail.com
// @host localhost:8080
// @BasePath /
func App() *fiber.App {

	//crear la instancia de la app

	app := fiber.New()

	// middleware 

	app.Use(cors.New())

	app.Use(recover.New())

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// swagger docs
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Crear routs de la aplicación

	router.Routers(app.Group("/api"))


	return app
}
