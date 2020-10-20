package application

import (
	infrastructure "github.com/nossey/northernlife/infrastructure"
	"golang.org/x/crypto/bcrypt"
)

// HashedPassword contains hashedPassword
type HashedPassword struct {
	HashedPassword string
}

// IsValidUser returns specified userId/password are valid or not
func IsValidUser(userID string, password string) bool {
	var hashedPassword HashedPassword
	sql := "select hashed_password from users where id = ?"
	// Spec:GORM's scan requires struct, not string
	infrastructure.Db.Raw(sql, userID).Scan(&hashedPassword)

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword.HashedPassword), []byte(password))
	return err == nil
}
