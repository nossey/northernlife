package infrastructure

import (
	"golang.org/x/crypto/bcrypt"
)

// ToHashedPassword returns hashed password
func ToHashedPassword(password string) (hash string, err error) {
	bytePassword := []byte(password)
	hashed, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	hash = string(hashed)

	return
}
