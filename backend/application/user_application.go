package application

import (
	"time"

	"github.com/nossey/northernlife/infrastructure"
)

// User is record form of users
type User struct {
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
	ID             string    `gorm:"column:id"`
	Email          string    `gorm:"column:email"`
	HashedPassword string    `gorm:"column:hashed_password"`
}

// CreateUser creates a user
func CreateUser(id string, email string, password string) {
	hashedPassword, _ := infrastructure.ToHashedPassword(password)
	user := User{CreatedAt: time.Now(), UpdatedAt: time.Now(), ID: id, Email: email, HashedPassword: hashedPassword}

	tx := infrastructure.Db.Begin()
	defer tx.Close()
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
}
