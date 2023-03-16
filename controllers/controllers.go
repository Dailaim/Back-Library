package controller

import (
/* 	"log"

	"github.com/Daizaikun/back-library/database"
	"gorm.io/gorm" */

	"github.com/gofiber/fiber/v2"
	
)

type Ctrl interface {
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
}




/* type CRUD struct{
	
	DB *gorm.DB
}

func NewCRUD(db *gorm.DB) *CRUD {
	return &CRUD{DB: db}
}


type Model interface{

	
}

func (a CRUD) GetById(id string, obj Model) ( error) {

	result := a.DB.First(obj, id)
	
	return result.Error
}

func (a CRUD) Create(obj Model) error {
	db :=database.Connect()

	log.Println(obj)
	result := db.Create(&obj)
	return result.Error
}

func (a CRUD) GetAll( obj Model) (error) {
	result := a.DB.Find(obj)
	return result.Error
}

func (a CRUD) Update(obj Model) error {
	result := a.DB.Save(obj)
	return result.Error
}

func (a CRUD) Delete(obj Model) error {
	result := a.DB.Delete(obj)
	return result.Error
} */

