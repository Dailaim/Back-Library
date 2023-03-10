package application

import (
	"github.com/Daizaikun/back-library/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/cors"
	
)

func App() *fiber.App {

	app := fiber.New()

	app.Use(cors.New())

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Route("/user", router.User)

	app.Route("/review", router.Review)
	
	app.Route("/book", router.Book)

	app.Static("/images/photos", "../uploads/photos")
	app.Static("/images/photos", "../uploads/ImagesBooks")

	return app
}
