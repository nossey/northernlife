package dataaccessor

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/nossey/northernlife/infrastructure"
)

// GetPostType is kind of get posts
type GetPostType int

const (
	// PublishedOnly get all published posts
	PublishedOnly GetPostType = iota
	// DraftOnly get all draft posts
	DraftOnly
	// All get all draft & published posts
	All
)

// PostListItem is single item of get posts
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
	Tags      pq.StringArray `gorm:"tags type:text[]"`
}

// PostDataAccessor provides access to posts for PostApplication
type PostDataAccessor struct {
}

// PostAccessor is the single object of PostDataAccessor
var PostAccessor *PostDataAccessor

func ini() {
	PostAccessor = &PostDataAccessor{}
}

// GetPostsCount get posts count with getPostType
func (accessor *PostDataAccessor) GetPostsCount(getPostType GetPostType) (count int) {
	filterSQL := ""
	switch getPostType {
	case PublishedOnly:
		filterSQL = `
where
	published = true
`
		break
	case DraftOnly:
		filterSQL = `
where
	published = false
`
		break
	case All:
		break
	}

	db := infrastructure.Db
	getPostsCountSQL := `
select
	count(1)
from
	posts
` + filterSQL
	row := db.Raw(getPostsCountSQL).Row()
	row.Scan(&count)
	return
}

// GetPosts get posts
func (accessor *PostDataAccessor) GetPosts(offset int, limit int, getPostType GetPostType) (result []PostListItem) {
	result = []PostListItem{}
	filterSQL := ""

	switch getPostType {
	case PublishedOnly:
		filterSQL = `
where
	published = true
`
		break
	case DraftOnly:
		filterSQL = `
where
	published = false 
`
		break
	case All:
		break
	}

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
		p.thumbnail,
		array_remove(array_agg(t.tag_name), null) as tags
	from
		posts p
	left join
		tags_posts_attachments tpa
		on p.id = tpa.post_id
	left join
		tags t
		on t.id = tpa.tag_id` +
		filterSQL +
		`group by
		p.id
	order by
		p.created_at desc
	offset
		?
	limit
		?
	`

	db := infrastructure.Db
	rows, err := db.Raw(getPostsSQL, offset, limit).Rows()
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var post PostListItem
		post.Tags = []string{}
		db.ScanRows(rows, &post)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, post)
	}
	return
}

// Post is posts table
type Post struct {
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
	ID        uuid.UUID `gorm:"id"`
	UserID    string    `gorm:"user_id"`
	Title     string    `gorm:"title"`
	Body      string    `gorm:"body"`
	PlainBody string    `gorm:"plain_body"`
	Published bool      `gorm:"published"`
	Thumbnail string    `gorm:"thumbnail"`
}
