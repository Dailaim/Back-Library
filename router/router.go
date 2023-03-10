package router

import (
	"github.com/Daizaikun/back-library/controllers"
	"github.com/Daizaikun/back-library/controllers/author"
	"github.com/Daizaikun/back-library/controllers/book"
	"github.com/Daizaikun/back-library/controllers/category"
	"github.com/Daizaikun/back-library/controllers/review"
	"github.com/Daizaikun/back-library/controllers/user"
	"github.com/gofiber/fiber/v2"
)

func CRUD(api fiber.Router, obj controller.Ctrl) {
	// Asignar el controlador a la ruta especificada
	api.Post("/", obj.Create)
	api.Get("/", obj.GetAll)
	api.Put("/:id", obj.Update)
	api.Delete("/:id", obj.Delete)
	api.Get("/:id", obj.Get)
}

func User(api fiber.Router) {

	CRUD(api, &user.Controller{})
	
}

func Book(api fiber.Router) {
	
	CRUD(api, &book.Controller{})

}

func Category(api fiber.Router) {
	
	CRUD(api, &category.Controller{})

}

func Review(api fiber.Router) {
	
	CRUD(api, &review.Controller{})

}

func Author(api fiber.Router) {
	
	CRUD(api, &author.Controller{})

}





