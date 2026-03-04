package security

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash generates a bcrypt hash from a plain text password.
// Returns the generated hash or an error if the process fails.
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// validatePassword compares a plain text password with a previously
// generated bcrypt hash. It returns nil if the password matches,
// or an error if it does not match or if the comparison fails.
func ValidatePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
