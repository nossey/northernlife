package model

import (
	"time"
)

// Post is a post
type Post struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ID        string    `json:"id"`
	UserID    string    `json:"userID"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	PlainBody string    `json:"plainBody"`
	Published bool      `json:"published"`
	Thumbnail string    `json:"thumbnail"`
	Tags      []string  `json:"tags"`
}

// PostListModel bundles Posts and TotalCount
type PostListModel struct {
	Posts        []Post `json:"posts"`
	TotalCount   int    `json:"totalCount"`
	PerPageCount int    `json:"perPageCount"`
}

// PostCreateBody is body to create a new post
type PostCreateBody struct {
	Body      string   `json:"body" binding:"required"`
	Title     string   `json:"title" binding:"required"`
	PlainBody string   `json:"plainBody" binding:"required"`
	Tags      []string `json:"tags" binding:"required"`
	Thumbnail string   `json:"thumbnail" binding:"required"`
}

// PostCreateResult contains results of post creation
type PostCreateResult struct {
	PostID string `json:"postID"`
}

// PostUpdateModel is body to update a post
type PostUpdateModel struct {
	Body      string   `json:"body" binding:"required"`
	Title     string   `json:"title" binding:"required"`
	PlainBody string   `json:"plainBody" binding:"required"`
	Tags      []string `json:"tags" binding:"required"`
	Thumbnail string   `json:"thumbnail" binding:"required"`
}

// PostUpdateResult is success result of update a post
type PostUpdateResult struct {
	PostID string `json:"postID"`
}

// PostDeleteResult contains results of post creation
type PostDeleteResult struct {
	PostID string `json:"postID"`
}
