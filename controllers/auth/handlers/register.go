package handlers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"github.com/Daizaikun/back-library/app/middleware"
	AuthModels "github.com/Daizaikun/back-library/controllers/auth/models"
	"github.com/Daizaikun/back-library/database"
	"github.com/Daizaikun/back-library/models"
)

// @Summary Register
// @Tags Auth
// @Description Register
// @Produce json
// @Param user body AuthModels.UserRegister true "User"
// @Success 200 {object} AuthModels.Response 
// @Failure 400	{object} AuthModels.Response
// @Failure 409 {object} AuthModels.Response
// @Router /auth/register [post]
func Registration(ctx *fiber.Ctx) error {
	var user AuthModels.UserRegister
	if err := ctx.BodyParser(&user); err != nil {
		return err
	}

	// Validar que el email y la contraseña no estén vacíos
	if user.Email == "" || user.Password == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.Error{
			Message: "El email y la contraseña son obligatorios",
			Code:    fiber.StatusBadRequest,
		})
	}

	// Validar que el email no esté en uso
	existingUser := models.User{}
	result := database.DB.Where("email = ?", user.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		return ctx.Status(fiber.StatusConflict).JSON(models.Error{
			Message: "El email ya está en uso",
			Code:    fiber.StatusConflict,
		})
	}

	// Hashear la contraseña del usuario
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(AuthModels.Response{
			Error: &AuthModels.Error{
				Message: "Error al encriptar la contraseña",
				Code:    fiber.StatusInternalServerError,
			},
			Data: nil,
		})
	}

	NewUser := models.User{
		Email:    user.Email,
		Password: string(hashedPassword),
		Photo:    user.Photo,
		Name:     user.Name,
	}

	// Guardar el usuario en la base de datos
	result = database.DB.Create(&NewUser)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(AuthModels.Response{
			Error: &AuthModels.Error{
				Message: "Error al guardar el usuario",
				Code:    fiber.StatusInternalServerError,
			},
			Data: nil,
		})
	}

	// Establecer el token de acceso en la estructura User
	tokenAccess, err := generateToken(NewUser.ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(AuthModels.Response{
			Error: &AuthModels.Error{
				Message: "Error al generar el token de acceso",
				Code:    fiber.StatusInternalServerError,
			},
			Data: nil,
		})
	}

	// Devolver la estructura User al cliente
	return ctx.Status(fiber.StatusOK).JSON(AuthModels.Response{
		Error: nil,
		Data: &AuthModels.Data{
			TokenAccess: tokenAccess,
		},
	},
	)
}

// Generar el token de acceso
func generateToken(id uint) (string, error) {

	claims := jwt.MapClaims{}
	claims["user_id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(middleware.SecretKey))

}
