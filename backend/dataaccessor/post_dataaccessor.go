package dataaccessor

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/nossey/northernlife/domain"
	"github.com/nossey/northernlife/infrastructure"
	"gorm.io/gorm"
)

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

// PostDataAccessor provides access to posts for PostApplication
type PostDataAccessor struct {
}

// PostAccessor is the single object of PostDataAccessor
var PostAccessor *PostDataAccessor

func ini() {
	PostAccessor = &PostDataAccessor{}
}

// GetPostsCount get posts count with getPostType
func (accessor *PostDataAccessor) GetPostsCount(getPostType domain.GetPostType) (count int) {
	filterSQL := ""
	switch getPostType {
	case domain.Published:
		filterSQL = `
where
	published = true
`
		break
	case domain.Draft:
		filterSQL = `
where
	published = false
`
		break
	case domain.All:
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

// GetAdminPostsCount get posts count
func (accessor *PostDataAccessor) GetAdminPostsCount(getPostType domain.GetPostType, userID string) (count int) {
	filterSQL := ""
	switch getPostType {
	case domain.Published:
		filterSQL = `
where
	published = true
	and user_id = ?
`
		break
	case domain.Draft:
		filterSQL = `
where
	published = false
	and user_id = ?
`
		break
	case domain.All:
		filterSQL = `
where
	user_id = ?
`
		break
	}

	db := infrastructure.Db
	getPostsCountSQL := `
select
	count(1)
from
	posts
` + filterSQL
	row := db.Raw(getPostsCountSQL, userID).Row()
	row.Scan(&count)
	return
}

// GetPosts get posts
func (accessor *PostDataAccessor) GetPosts(offset int, limit int, getPostType domain.GetPostType) (result []domain.PostListItem) {
	result = []domain.PostListItem{}
	filterSQL := ""

	switch getPostType {
	case domain.Published:
		filterSQL = `
where
	published = true
`
		break
	case domain.Draft:
		filterSQL = `
where
	published = false 
`
		break
	case domain.All:
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
		var post domain.PostListItem
		post.Tags = []string{}
		db.ScanRows(rows, &post)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, post)
	}
	return
}

// GetPostList get posts
func (accessor *PostDataAccessor) GetPostList(tags []string, offset int, limit int, searchWord string, getPostType domain.GetPostType) (result []domain.PostListItem) {
	filterSQL := ""
	switch getPostType {
	case domain.Published:
		filterSQL = `
where
	pwt.published = true
	and pwt.tags @> ?
	and pwt.plain_body ilike ?`
		break
	case domain.Draft:
		filterSQL = `
where
	pwt.published = false 
	and pwt.tags @> ?
	and pwt.plain_body ilike ?`
		break
	case domain.All:
		filterSQL = `
where
	pwt.tags @> ?
	and pwt.plain_body ilike ?`
		break
	}

	getPostsSQL := `
with posts_with_tags as (
select
	p.created_at as created_at,
	p.updated_at as updated_at,
	p.id as id,
	p.user_id as user_id,
	p.title as title,
	p.body as body,
	p.plain_body as plain_body,
	p.published as published,
	p.thumbnail as thumbnail,
	array_remove(array_agg(t.tag_name), null) as tags
from
	posts p
left join
	tags_posts_attachments tpa
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
)
select
	*
from
	posts_with_tags pwt` + filterSQL

	db := infrastructure.Db
	tagsSpecified := pq.StringArray(tags)
	search := "%" + searchWord + "%"
	rows, err := db.Raw(getPostsSQL, offset, limit, tagsSpecified, search).Rows()
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var post domain.PostListItem
		post.Tags = []string{}
		db.ScanRows(rows, &post)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, post)
	}
	return
}

// GetPostListCount get count of posts
func (accessor *PostDataAccessor) GetPostListCount(tags []string, getPostType domain.GetPostType) (count int) {
	filterSQL := ""
	switch getPostType {
	case domain.Published:
		filterSQL = `
where
	pwt.published = true
	and pwt.tags @> ?`
		break
	case domain.Draft:
		filterSQL = `
where
	pwt.published = false 
	and pwt.tags @> ?`
		break
	case domain.All:
		filterSQL = `
where
	pwt.tags @> ?`
		break
	}

	getCountSQL := `
with posts_with_tags as (
select
	p.user_id as user_id,
	p.title as title,
	p.body as body,
	p.plain_body as plain_body,
	p.published as published,
	array_remove(array_agg(t.tag_name), null) as tags
from
	posts p
left join
	tags_posts_attachments tpa
	on p.id = tpa.post_id
left join
	tags t
	on t.id = tpa.tag_id
group by
	p.id
order by
	p.created_at desc
)
select
	count(1)
from
	posts_with_tags pwt` + filterSQL

	db := infrastructure.Db
	tagsSpecified := pq.StringArray(tags)
	row := db.Raw(getCountSQL, tagsSpecified).Row()
	row.Scan(&count)
	return
}

// GetAdminPosts get posts
func (accessor *PostDataAccessor) GetAdminPosts(offset int, limit int, getPostType domain.GetPostType, userID string) (result []domain.PostListItem) {
	result = []domain.PostListItem{}
	filterSQL := ""

	switch getPostType {
	case domain.Published:
		filterSQL = `
where
	p.published = true
	and p.user_id = ?
`
		break
	case domain.Draft:
		filterSQL = `
where
	p.published = false 
	and p.user_id = ?
`
		break
	case domain.All:
		filterSQL = `
where
	p.user_id = ?
`
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
	rows, err := db.Raw(getPostsSQL, userID, offset, limit).Rows()
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var post domain.PostListItem
		post.Tags = []string{}
		err = db.ScanRows(rows, &post)
		if err != nil {
			fmt.Println(err)
		} else {
			result = append(result, post)
		}
	}
	return

}

// CreatePost create a post
func (accessor *PostDataAccessor) CreatePost(p Post, tx *gorm.DB) (err error) {
	err = tx.Create(&p).Error
	return
}

// CreateTagAttachments attach tags to post
func (accessor *PostDataAccessor) CreateTagAttachments(postID string, tags pq.StringArray, tx *gorm.DB) (err error) {
	tagsAttachmentSQL := `
insert into tags_posts_attachments
(
	created_at,
	id,
	post_id,
	tag_id
)
select
	current_timestamp as created_at,
	uuid_generate_v4() as id,
	? as post_id,
	t.id as tag_id
from
	tags t
where
	tag_name = any(?)
`
	err = tx.Exec(tagsAttachmentSQL, postID, tags).Error
	return
}

// UpdatePost update a post
func (accessor *PostDataAccessor) UpdatePost(postID uuid.UUID, title string, body string, plainBody string, published bool, thumbnail string, userID string, tx *gorm.DB) (err error) {
	updatePostSQL := `
update posts 
set 
	updated_at = current_timestamp,
	title = ?,
	body = ?,
	plain_body = ?,
	published = ?,
	thumbnail = ? 
where
	id = ? 
	and user_id = ?;`

	err = tx.Exec(updatePostSQL, title, body, plainBody, published, thumbnail, postID, userID).Error
	return
}

// DeleteAttachedTags delete tags attached to post
func (accessor *PostDataAccessor) DeleteAttachedTags(postID uuid.UUID, tx *gorm.DB) (err error) {
	deleteAttachedTagsSQL := `
delete from
	tags_posts_attachments
where
	post_id = ?`

	err = tx.Exec(deleteAttachedTagsSQL, postID).Error
	return
}

// DeletePost delete a post
func (accessor *PostDataAccessor) DeletePost(postID uuid.UUID, tx *gorm.DB) (err error) {
	deletePostSQL := `
delete from
	posts
where
	id = ?	
	`
	result := tx.Exec(deletePostSQL, postID)
	if result.RowsAffected != 1 {
		err = errors.New("No post deleted")
	}

	return
}
