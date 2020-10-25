package application

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/nossey/northernlife/infrastructure"
	"github.com/nossey/northernlife/model"
)

// GetPosts get posts with pagination
func GetPosts(page int) (result model.PostListModel) {
	perPageCount := 10
	offset := perPageCount * (page - 1)
	result.Posts = []model.Post{}
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
		var post model.Post
		db.ScanRows(rows, &post)
		if err != nil {
			fmt.Println(err)
		}
		result.Posts = append(result.Posts, post)
	}
	return
}

// GetPost get post with specified id
func GetPost(id uuid.UUID) (post model.Post, err error) {
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