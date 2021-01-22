package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

// GetPostType is kind of get posts
type GetPostType int

const (
	// Published get all published posts
	Published GetPostType = iota
	// Draft get all draft posts
	Draft
	// All get all draft & published posts
	All
)

// PostListResult is domain layer's post list result
type PostListResult struct {
	Posts        []PostListItem
	TotalCount   int
	PerPageCount int
}

// PostListItem is single item when get posts
type PostListItem struct {
	CreatedAt time.Time      `gorm:"created_at"`
	UpdatedAt time.Time      `gorm:"updated_at"`
	ID        uuid.UUID      `gorm:"id"`
	UserID    string         `gorm:"user_id"`
	Title     string         `gorm:"title"`
	Body      string         `gorm:"body"`
	PlainBody string         `gorm:"plain_body"`
	Published bool           `gorm:"published"`
	Thumbnail string         `gorm:"thumbnail"`
	Tags      pq.StringArray `gorm:"type:text[]"`
}

// SinglePostItem is the result of get single post
type SinglePostItem struct {
	CreatedAt time.Time      `gorm:"created_at"`
	UpdatedAt time.Time      `gorm:"updated_at"`
	ID        uuid.UUID      `gorm:"id"`
	UserID    string         `gorm:"user_id"`
	Title     string         `gorm:"title"`
	Body      string         `gorm:"body"`
	PlainBody string         `gorm:"plain_body"`
	Published bool           `gorm:"published"`
	Thumbnail string         `gorm:"thumbnail"`
	Tags      pq.StringArray `gorm:"tags type:text[]"`
}
