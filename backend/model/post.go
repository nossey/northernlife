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
	Thumbnail string    `json:"thumbnail"`
	Tags      []string  `json:"tags"`
}

// PostListModel bundles Posts and TotalCount
type PostListModel struct {
	Posts        []Post `json:"posts"`
	TotalCount   int    `json:"total_count"`
	PerPageCount int    `json:"per_page_count"`
}

// PostCreateBody is body to create a new post
type PostCreateBody struct {
	Body      string   `json:"body" binding:"required"`
	Title     string   `json:"title" binding:"required"`
	PlainBody string   `json:"plain_body" binding:"required"`
	Tags      []string `json:"tags" binding:"required"`
	Thumbnail string   `json:"thumbnail" binding:"required"`
}

// PostCreateResult contains results of post creation
type PostCreateResult struct {
	PostID string `json:"post_id"`
}
