package auth

import (
	"fmt"
	"mime/multipart"
	"time"

	userCRUD "github.com/Daizaikun/back-library/controllers/user"
	"github.com/Daizaikun/back-library/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Photo    *multipart.FileHeader `json:"photo"`
}

func Register(c *fiber.Ctx) error {
	var input RegisterInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	// Verificar si el email ya está registrado
	user := new(models.User)
	if err := userCRUD.GetUserByEmail(user, input.Email); err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Email already exists",
		})
	}

	// Hash de la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to hash password",
		})
	}

	// Guardar la imagen en la carpeta de fotos de usuario
	photoID := uuid.New().String()
	if err := c.SaveFile(input.Photo, fmt.Sprintf("./uploads/photos/%s", photoID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to save photo",
		})
	}

	// Crear el nuevo usuario
	newUser := &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Photo:    photoID,
	}
	if err := userCRUD.CreateUser(newUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create user",
		})
	}

	// Crear el token JWT
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = newUser.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // El token expira en 24 horas

	tokenString, err := token.SignedString([]byte("library"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create token",
		})
	}

	// Devolver la respuesta JSON con el objeto creado y el token JWT
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user":  newUser,
		"token": tokenString,
	})
}