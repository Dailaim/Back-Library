package user

import (
	"github.com/Daizaikun/back-library/app/middleware"
	"github.com/Daizaikun/back-library/controllers/user/handlers"
	"github.com/gofiber/fiber/v2"
)


func Routers(api fiber.Router) {

	api.Get("/", handlers.GetAll)

	// Proteger todas las rutas de usuario con el middleware de autenticación

	api.Use(middleware.AuthMiddleware)
	
	// Definir las rutas de usuario
	api.Get("/:id", handlers.GetById)
	api.Put("/", handlers.Update)
	api.Delete("/", handlers.Delete)

}