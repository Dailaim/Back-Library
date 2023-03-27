package router

import (
	"github.com/gofiber/fiber/v2"

	controller "github.com/Daizaikun/back-library/controllers"
	"github.com/Daizaikun/back-library/controllers/author"
	"github.com/Daizaikun/back-library/controllers/book"
	"github.com/Daizaikun/back-library/controllers/category"

)

func CRUD(app fiber.Router) {

	app.Route("/user", User)

	app.Route("/review", Review)

	app.Route("/book", Book)

	app.Route("/author", Author)

	app.Route("/category", Category)
	
}


// Esto es un ejemplo de como se puede hacer un CRUD con un controlador
func BasicCRUD(api fiber.Router, obj controller.Ctrl) {
	// Asignar el controlador a la ruta especificada
	api.Post("/", obj.Create)
	api.Get("/", obj.GetAll)
	api.Put("/:id", obj.Update)
	api.Delete("/:id", obj.Delete)
	api.Get("/:id", obj.GetById)

}

// @Summary Get all books
// @Tags Books
// @Description Get all books
// @Produce json
// @Success 200 {array} Book
// @Router /books [get]
func Book(api fiber.Router) {

	BasicCRUD(api, &book.Controller{})

}

func Category(api fiber.Router) {

	BasicCRUD(api, &category.Controller{})

}


func Author(api fiber.Router) {

	BasicCRUD(api, &author.Controller{})

}
