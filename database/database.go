package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	/* 	"gorm.io/gorm/logger" */)

var DB *gorm.DB



func Connect() (*gorm.DB) {

	host := env("host_db")

	user := env("user_db")

	password := env("password_db")

	name := env("name_db")

	port := env("port_db")

	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%s
	sslmode=disable TimeZone=America/Bogota`, host, user, password, name, port)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

/*   DB.Logger = logger.Default.LogMode(logger.Info) */

  Migrate(DB)

  if err != nil {
		panic("failed to connect database")
	}


  fmt.Println("Connected Successfully to the Database")

  return DB

}

func env(env string) string {
	variable := os.Getenv(env)
	if len(variable) == 0 {
		panic(fmt.Sprintf(`failed variable de entorno %s`, env))
	}
	return variable
}

/*
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)
*/
