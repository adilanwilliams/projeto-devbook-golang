package authentication

import (
	"devbook/src/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// CreateToken generates a JWT token for the given user ID.
func CreateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{
		"authorized": true,
		"exp":        time.Now().Add(time.Hour).Unix(),
		"userID":     userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	secretKey := config.SecretKey

	return token.SignedString(secretKey)
}

func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, getVerificationKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Token invalid.")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	tokenSplited := strings.Split(token, " ")

	if len(tokenSplited) == 2 {
		return tokenSplited[1]
	}

	return ""
}

func getVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Asing method unspered %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
