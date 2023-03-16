package router

import (
	"github.com/Daizaikun/back-library/controllers/auth"
	"github.com/gofiber/fiber/v2"
)

func Auth(app fiber.Router) {

	app.Post("/login", auth.Login)
	app.Post("/register", auth.Register)
	app.Post("Post", auth.Logout)

}
