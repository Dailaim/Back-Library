package main

import (
	"os"

	"github.com/Daizaikun/back-library/app"
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/helpers"
)

func main() {

	//Inicia la conexión a la base de datos

	db := database.Connect()
	database.Migrate(db)

	//Asegura que las carpetas existan

	helpers.DirSafe("./uploads")

	helpers.DirSafe("./uploads/ImagesBooks")

	helpers.DirSafe("./uploads/photos")

	// Activa el funcionamiento de la app

	app := application.App()

	app.Listen(port("3000"))
}

func port(p string) string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = p
	}
	return ":" + port
}
