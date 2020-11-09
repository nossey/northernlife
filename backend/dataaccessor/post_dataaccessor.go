package dataaccessor

import (
	"time"

	"github.com/google/uuid"
)

// Post is posts table
type Post struct {
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
	ID        uuid.UUID `gorm:"id"`
	UserID    string    `gorm:"user_id"`
	Title     string    `gorm:"title"`
	Body      string    `gorm:"body"`
	PlainBody string    `gorm:"plain_body"`
	Published bool      `gorm:"published"`
}
