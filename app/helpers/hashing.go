package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hash the password
func HashPassword(userPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// CheckPasswordHash compare the hash of the password with the password used by the user
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
