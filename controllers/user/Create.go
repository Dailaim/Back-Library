package user

import "github.com/gofiber/fiber/v2"

func (c *Controller) Create(ctx *fiber.Ctx) error{
	return ctx.SendString("Hola mundo")
}