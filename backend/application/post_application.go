package application

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nossey/northernlife/infrastructure"
)

// Post is application layer's post
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

// PostListResult is application layer's post list result
type PostListResult struct {
	Posts      []Post
	TotalCount int
}

// GetPosts get posts with pagination
func GetPosts(page int) (result PostListResult) {
	perPageCount := 10
	offset := perPageCount * (page - 1)
	result.Posts = []Post{}
	db := infrastructure.Db
	getPostsCountSQL := `
select
	count(1)
from
	posts
`
	row := db.Raw(getPostsCountSQL).Row()
	row.Scan(&result.TotalCount)

	getPostsSQL := `
select
	created_at,
	updated_at,
	id,
	user_id,
	title,
	body,
	plain_body,
	published
from
	posts
offset
	?
limit
	?
`
	rows, err := db.Raw(getPostsSQL, offset, perPageCount).Rows()
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var post Post
		db.ScanRows(rows, &post)
		if err != nil {
			fmt.Println(err)
		}
		result.Posts = append(result.Posts, post)
	}
	return
}

// GetPost get post with specified id
func GetPost(id uuid.UUID) (post Post, err error) {
	db := infrastructure.Db
	sql := `
select
	created_at,
	updated_at,
	id,
	user_id,
	title,
	body,
	plain_body,
	published
from
	posts
where
	id = ?
`
	err = db.Raw(sql, id.String()).First(&post).Error
	return
}
