package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {

	dsn := "host=localhost user=library password=controlbox dbname=library port=5432 sslmode=disable TimeZone=America/Bogota"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
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