package dataaccessor

import (
	"time"

	"github.com/google/uuid"
)

// Tag is column for table tags
type Tag struct {
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
	ID        uuid.UUID `gorm:"id"`
	TagName   string    `gorm:"tag_name"`
	UserID    string    `gorm:"user_id"`
}
