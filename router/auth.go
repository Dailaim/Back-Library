package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Daizaikun/back-library/app/middleware"
	"github.com/Daizaikun/back-library/controllers/auth"
	"github.com/Daizaikun/back-library/controllers/review"
	"github.com/Daizaikun/back-library/controllers/user"

)

func Auth(app fiber.Router) {

	app.Post("/login", auth.HandleAuthentication)
	app.Post("/register", auth.HandleRegistration)
	app.Post("/logout", auth.HandleLogout)

}

func User(api fiber.Router) {

	userCurd := &user.Controller{}
	// Proteger todas las rutas de usuario con el middleware de autenticación

	api.Use(middleware.AuthMiddleware)
	

	// Definir las rutas de usuario
	api.Get("/:id", userCurd.GetById)
	api.Put("/", userCurd.Update)
	api.Delete("/", userCurd.Delete)

}

func Review(api fiber.Router) {

	reviewCrud := &review.Controller{}

	api.Post("/book/:id", review.GetAllByBook)
	api.Get("/", reviewCrud.GetAll)

	api.Use(middleware.AuthMiddleware)

	api.Post("/", reviewCrud.Create)
	api.Put("/:id", reviewCrud.Update)
	api.Delete("/:id", reviewCrud.Delete)
	api.Get("/:id", reviewCrud.GetById)

}
