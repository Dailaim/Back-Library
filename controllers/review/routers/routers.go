package review

import (
	"github.com/Daizaikun/back-library/app/middleware"
	"github.com/Daizaikun/back-library/controllers/review/handlers"

	"github.com/gofiber/fiber/v2"
)

func Routers(api fiber.Router) {

	api.Get("/book/:id", handlers.GetAllByBook)
	api.Get("/", handlers.GetAll)

	api.Use(middleware.AuthMiddleware)

	api.Post("/", handlers.Create)
	api.Put("/:id", handlers.Update)
	api.Delete("/:id", handlers.Delete)
	api.Get("/:id", handlers.GetById)

}