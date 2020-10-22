package application

import (
	"fmt"

	"github.com/nossey/northernlife/infrastructure"
	"github.com/nossey/northernlife/model"
)

// GetPosts get posts with pagination
func GetPosts(page int) (result model.PostListModel) {
	db := infrastructure.Db
	result.Posts = []model.Post{}
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
`
	rows, err := db.Raw(getPostsSQL).Rows()
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
