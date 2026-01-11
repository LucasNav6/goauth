package utilities

import (
	"time"

	goauth_models "github.com/LucasNav6/goauth/pkg/models"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(config goauth_models.Configuration, userID string) string {
	// Define the secret key used for signing the token
	secretKey := []byte(config.Secret)

	// Create a new token object
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Second * time.Duration(config.SessionDurationInSeconds)).Unix(),
		"iat": time.Now().Unix(),
		"sub": userID,
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return ""
	}

	return tokenString
}
