package book

import (
	"github.com/Daizaikun/back-library/controllers/book/handlers"
	"github.com/gofiber/fiber/v2"
)

func Routers(app fiber.Router) {
	app.Get("/", handlers.GetAll)
	app.Get("/:id", handlers.GetById)
	app.Post("/", handlers.Create)
	app.Put("/:id", handlers.Update)
	app.Delete("/:id", handlers.Delete)
}