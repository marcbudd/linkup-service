package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWTToken(userID uint) (string, error) {
	// Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return tokenString, err
}
