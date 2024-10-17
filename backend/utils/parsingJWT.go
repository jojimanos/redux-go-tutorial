package utils

import (
	"errors"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

// Claims defines the structure for the JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Initialize your JWT secret
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func ParseJWT(tokenStr string) (*Claims, error) {
	// Parse the token using the secret key
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure that the token is signed using the HMAC method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		log.Println("Error parsing token:", err)
		return nil, err
	}

	// Validate the token and extract the claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	// If the token is invalid or claims cannot be extracted
	return nil, errors.New("invalid token")
}
