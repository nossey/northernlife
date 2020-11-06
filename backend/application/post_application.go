package application

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/nossey/northernlife/infrastructure"
)

// Post is application layer's post
type Post struct {
	CreatedAt time.Time      `gorm:"created_at"`
	UpdatedAt time.Time      `gorm:"updated_at"`
	ID        uuid.UUID      `gorm:"id"`
	UserID    string         `gorm:"user_id"`
	Title     string         `gorm:"title"`
	Body      string         `gorm:"body"`
	PlainBody string         `gorm:"plain_body"`
	Published bool           `gorm:"published"`
	Tags      pq.StringArray `gorm:"tags type:text[]"`
}

// PostListResult is application layer's post list result
type PostListResult struct {
	Posts        []Post
	TotalCount   int
	PerPageCount int
}

// GetPosts get posts with pagination
func GetPosts(page int) (result PostListResult) {
	perPageCount := 10
	result.PerPageCount = perPageCount
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
	p.created_at,
	p.updated_at,
	p.id,
	p.user_id,
	p.title,
	p.body,
	p.plain_body,
	p.published,
	coalesce(array_agg(t.tag_name) filter (where t.tag_name is not null), '{}') as tags
from
	posts p
left join
	tags_posts_attachment tpa
	on p.id = tpa.post_id
left join
	tags t
	on t.id = tpa.tag_id
group by
	p.id
order by
	p.created_at desc
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
		post.Tags = []string{}
		db.ScanRows(rows, &post)
		//rows.Scan(&post.CreatedAt, &post.UpdatedAt, &post.ID, &post.UserID, &post.Title, &post.Body, &post.PlainBody, &post.Published, pq.Array(&post.Tags))
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
