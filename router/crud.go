package router

import (
	controller "github.com/Daizaikun/back-library/controllers"
	"github.com/Daizaikun/back-library/controllers/author"
	"github.com/Daizaikun/back-library/controllers/book"
	"github.com/Daizaikun/back-library/controllers/category"

	"github.com/gofiber/fiber/v2"
)

func CRUD(app fiber.Router) {

	app.Route("/review", Review)

	app.Route("/book", Book)

	app.Route("/author", Author)

	app.Route("/category", Category)

	app.Route("/user", User)

}

func BasicCRUD(api fiber.Router, obj controller.Ctrl) {
	// Asignar el controlador a la ruta especificada
	api.Post("/", obj.Create)
	api.Get("/", obj.GetAll)
	api.Put("/:id", obj.Update)
	api.Delete("/:id", obj.Delete)
	api.Get("/:id", obj.GetById)
}

func Book(api fiber.Router) {

	BasicCRUD(api, &book.Controller{})

}

func Category(api fiber.Router) {

	BasicCRUD(api, &category.Controller{})

}

func Author(api fiber.Router) {

	BasicCRUD(api, &author.Controller{})

}
