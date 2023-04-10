package handlers

import (
	"fmt"

	"github.com/Daizaikun/back-library/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// UploadImagePhoto godoc
// @Summary Upload image photo
// @Description Upload images user photo
// @Tags images
// @Accept mpfd
// @Produce  json
// @Param image formData file true "image"
// @Success 200 {object} models.Image
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /images/upload/photo [post]
func UploadImagePhoto(ctx *fiber.Ctx) error {

	image := new(models.Image)

	file, err := ctx.FormFile("image")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.Error{
			Code:    fiber.StatusBadRequest,
			Message: "Error al subir la imagen",
			Error:   err,
		})
	}

	image.Image = uuid.New().String()

	err = ctx.SaveFile(file, fmt.Sprintf("./uploads/photos/%s", image.Image))

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Error{
			Code:    fiber.StatusInternalServerError,
			Message: "Error al guardar la imagen",
			Error:   err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(image)

}
