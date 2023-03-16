package application

import (
	"github.com/Daizaikun/back-library/router"
	"github.com/gofiber/fiber/v2"
/* 	jwtware "github.com/gofiber/jwt/v3" */
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/cors"
	
)

func App() *fiber.App {

	//crear la instancia de la app

	app := fiber.New()

	//descifra los datos

/* 	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("library"),
	})) */

	// middleware

	app.Use(cors.New())

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// Crear rout's de la aplicación 
	
	router.CRUD(app.Group("/crud"))

	
	router.Auth(app.Group("/auth"))

	//sirve las images de la aplicación


	router.Image(app.Group("/images/"))

	return app
}
