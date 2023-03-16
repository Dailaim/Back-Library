package images

import (
	"fmt"

	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UploadImageBooks(ctx *fiber.Ctx)error{

	image := new(models.Image)
	
	file, err := ctx.FormFile("image")

	if err != nil {
		return err
	}

	image.Image = uuid.New().String()

	err = ctx.SaveFile(file, fmt.Sprintf("./uploads/ImagesBooks/%s", image.Image))

	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(image)

}


func UploadImagePhoto(ctx *fiber.Ctx)error{

	image := new(models.Image)
	
	file, err := ctx.FormFile("image")

	if err != nil {
		return err
	}

	image.Image = uuid.New().String()

	err = ctx.SaveFile(file, fmt.Sprintf("./uploads/photos/%s", image.Image))

	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(image)

}