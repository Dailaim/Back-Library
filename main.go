package main

import (
	"github.com/joho/godotenv"

	"github.com/Daizaikun/back-library/app"
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/helpers"

)


func main() {

	//inicia variables de entorno si existen
	godotenv.Load()

	//Inicia la conexión a la base de datos

	database.DB = database.Connect()

	//Asegura que las carpetas existan

	helpers.DirSafe("./uploads")

	helpers.DirSafe("./uploads/ImagesBooks")

	helpers.DirSafe("./uploads/photos")

	// Activa el funcionamiento de la app

	app := app.App()

	app.Listen(":" + helpers.Env("port","8080"))
}


