package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// Clave secreta para firmar tokens JWT
const JWTSecret = "library"

// Tiempo de expiración de los tokens JWT (24 horas)
const JWTExpiration = time.Hour * 24

type Claims struct {
	jwt.StandardClaims
	Email string `json:"email"`
}

type AuthController struct{}