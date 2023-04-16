package images

import (
	"github.com/Daizaikun/back-library/controllers/images/handlers"
	"github.com/gofiber/fiber/v2"
)



func Routers(app fiber.Router) {
	
	app.Static("/photos/", handlers.UsersPhotos())

	app.Static("/books/", handlers.BooksImages())

	app.Post("/upload/book", handlers.UploadImageBooks)

	app.Post("/upload/photo", handlers.UploadImagePhoto)
}


