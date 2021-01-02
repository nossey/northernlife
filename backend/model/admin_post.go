package model

import (
	"time"
)

// AdminPostListItem is a post
type AdminPostListItem struct {
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

// AdminPostListModel bundles Posts and TotalCount
type AdminPostListModel struct {
	Posts        []AdminPostListItem `json:"posts"`
	TotalCount   int                 `json:"totalCount"`
	PerPageCount int                 `json:"perPageCount"`
}
