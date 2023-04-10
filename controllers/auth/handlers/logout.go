package handlers

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"

	"github.com/Daizaikun/back-library/app/middleware"
	"github.com/Daizaikun/back-library/models"
)

// HandleLogout godoc
// @Summary Logout
// @Description Logout
// @Tags Auth
// @Accept json
// @Produce json
// @Success 204 
// @Failure 401 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /auth/logout [post]
func Logout(ctx *fiber.Ctx) error {
	// Obtener el token de acceso del encabezado Authorization
	authHeader := ctx.Get("Authorization")
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	// Invalidar el token de acceso
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(middleware.SecretKey), nil
	})

	if err != nil || !token.Valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(models.Error{
			Message: "Token inválido",
			Code:    fiber.StatusUnauthorized,
		})
	}

	// Agregar el token a la lista negra para invalidarlo
	err = middleware.BlackListAddToken(tokenString)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Error{
			Message: "Hubo un error al invalidar el token",
			Code:    fiber.StatusInternalServerError,
		})
	}

	// Devolver una respuesta vacía con un estado HTTP 204 No Content
	return ctx.SendStatus(fiber.StatusNoContent)
}
