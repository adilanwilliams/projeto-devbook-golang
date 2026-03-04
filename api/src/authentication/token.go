package authentication

import (
	"devbook/src/config"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// CreateToken generates a JWT token for the given user ID.
func CreateToken(userID uint64) (string, error){
	permissions := jwt.MapClaims{
		"authorized": true,
		"exp": time.Now().Add(time.Hour).Unix(),
		"userID": userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	secretKey := config.SecretKey

	return token.SignedString(secretKey)
}