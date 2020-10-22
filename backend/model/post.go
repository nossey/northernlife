package model

import (
	"time"

	"github.com/google/uuid"
)

// Post is a post
type Post struct {
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
	ID        uuid.UUID `json:"id" gorm:"id"`
	UserID    string    `json:"user_id" gorm:"user_id"`
	Title     string    `json:"title" gorm:"title"`
	Body      string    `json:"body" gorm:"body"`
	PlainBody string    `json:"plain_body" gorm:"plain_body"`
	Published bool      `json:"published" gorm:"published"`
}

// PostListModel bundles Posts and TotalCount
type PostListModel struct {
	Posts      []Post `json:"posts"`
	TotalCount int    `json:"total_count"`
}
