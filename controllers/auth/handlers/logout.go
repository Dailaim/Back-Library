package handlers

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/Daizaikun/back-library/app/middleware"
	AuthModels "github.com/Daizaikun/back-library/controllers/auth/models"
	"github.com/Daizaikun/back-library/helpers"
)

// HandleLogout godoc
// @Summary Logout
// @Description Logout
// @Tags Auth
// @Accept json
// @Produce json
// @Success 204
// @Failure 401 {object} AuthModels.Response
// @Failure 500 {object} AuthModels.Response
// @Router /auth/logout [post]
func Logout(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")

	if !strings.HasPrefix(authHeader, "Bearer") {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization header must be in the format of Bearer [token]",
		})
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	if _, err := helpers.ValidateToken(tokenString); err != nil {
		log.Printf("JWT Token validation. %v", err)

		return ctx.Status(fiber.StatusUnauthorized).JSON(AuthModels.Response{
			Data: nil,
			Error: &AuthModels.Error{
				Message: "Invalid access token",
				Code:    fiber.StatusUnauthorized,
			},
		},
		)
	}

	// Add the token to the blacklist to invalidate it
	if err := middleware.BlackListAddToken(tokenString); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(AuthModels.Response{
			Data: nil,
			Error: &AuthModels.Error{
				Message: "Could not invalidate access token",
				Code:    fiber.StatusInternalServerError,
			},
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
