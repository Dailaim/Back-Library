package helpers

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SecretKey = "library"

// ValidateToken validates a JSON Web Token (JWT) string and returns its claims.
func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure the signing method is correct.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(SecretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); !ok {
		return nil, fmt.Errorf("failed to extract claims from token")
	} else {
		return claims, nil
	}
}

// GenerateToken creates a new access token for a given user ID using the HS256 signing method.
func GenerateToken(id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(SecretKey))
}
