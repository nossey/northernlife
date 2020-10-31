package model

import (
	"time"
)

// Post is a post
type Post struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	PlainBody string    `json:"plain_body"`
	Published bool      `json:"published"`
}

// PostListModel bundles Posts and TotalCount
type PostListModel struct {
	Posts      []Post `json:"posts"`
	TotalCount int    `json:"total_count"`
}
