package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	/* 	"gorm.io/gorm/logger" */

	"github.com/Daizaikun/back-library/helpers"
)

var DB *gorm.DB

// Inicia la conexión a la base de datos
func Connect() *gorm.DB {

	//Comprueba todas las variables de entorno6

	host := helpers.Env("host_db", "localhost")

	user := helpers.Env("user_db", "library")

	password := helpers.Env("password_db", "controlbox")

	name := helpers.Env("name_db", "library")

	port := helpers.Env("port_db", "5432")

	// Ruta para la conexión de postgres

	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%s
	sslmode=disable TimeZone=America/Bogota`, host, user, password, name, port)

	// Conexión a la base de datos

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	

	/* DB.Logger = logger.Default.LogMode(logger.Info) */

	err =Migrate(DB)
	if err != nil {
		panic("failed to migrate database")
	}

	return DB

}
