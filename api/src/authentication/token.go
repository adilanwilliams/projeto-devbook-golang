package authentication

import (
	"devbook/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// UserIDToken defines the key used to store the user ID inside the JWT claims.
var UserIDToken = "userID"

// CreateToken generates a JWT token containing the user ID and authorization data.
// The token is signed using the application's secret key and has an expiration time.
func CreateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{
		"authorized": true,
		"exp":        time.Now().Add(time.Hour).Unix(),
		UserIDToken:  userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	secretKey := config.SecretKey

	return token.SignedString(secretKey)
}

// ValidateToken verifies if the JWT token present in the request is valid.
// It checks the token signature and ensures the token has not expired.
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

// ExtractUserID retrieves the user ID stored inside the JWT token
// present in the request Authorization header.
func ExtractUserID(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, getVerificationKey)

	if err != nil {
		return 0, err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		permissionsUserID := fmt.Sprintf("%.0f", permissions[UserIDToken])
		userID, err := strconv.ParseUint(permissionsUserID, 10, 64)

		if err != nil {
			return 0, err
		}

		return userID, nil
	}

	return 0, errors.New("Invalid Token")
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
