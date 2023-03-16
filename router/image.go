package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Daizaikun/back-library/controllers/images"
)

func Image(app fiber.Router) {

	app.Static("/photos/", "./uploads/photos")

	app.Static("/books/", "./uploads/ImagesBooks")

	app.Post("/upload/book", images.UploadImageBooks)

	app.Post("/upload/photo", images.UploadImagePhoto)

}