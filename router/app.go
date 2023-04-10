package router

import (

	"github.com/Daizaikun/back-library/controllers/auth/routers"
	"github.com/Daizaikun/back-library/controllers/author/routers"
	"github.com/Daizaikun/back-library/controllers/book/routers"
	"github.com/Daizaikun/back-library/controllers/category/routers"
	"github.com/Daizaikun/back-library/controllers/images/routers"
	"github.com/Daizaikun/back-library/controllers/review/routers"
	"github.com/Daizaikun/back-library/controllers/user/routers"
	"github.com/gofiber/fiber/v2"
)


func Routers(app fiber.Router) {

	auth.Routers(app.Group("/auth"))

	images.Routers(app.Group("/images"))

	crud := app.Group("/book")
	
	book.Routers(crud.Group("/book"))
	author.Routers(crud.Group("/author"))
	category.Routers(crud.Group("/category"))
	review.Routers(crud.Group("/review"))
	user.Routers(crud.Group("/user"))
	
}
