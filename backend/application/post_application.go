package application

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"github.com/nossey/northernlife/dataaccessor"
	"github.com/nossey/northernlife/domain"
	"github.com/nossey/northernlife/infrastructure"
)

// PostApplication manages posts
type PostApplication struct {
}

// PostApp provides functions for post
var PostApp *PostApplication

func init() {
	PostApp = &PostApplication{}
}

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
	Thumbnail string         `gorm:"thumbnail"`
	Tags      pq.StringArray `gorm:"tags type:text[]"`
}

// PostCreate is used to create a post
type PostCreate struct {
	UserID    string
	Title     string
	Body      string
	PlainBody string
	Published bool
	Thumbnail string
	Tags      pq.StringArray
}

// PostUpdate is used to update a post
type PostUpdate struct {
	PostID    uuid.UUID
	UserID    string
	Title     string
	Body      string
	PlainBody string
	Published bool
	Thumbnail string
	Tags      pq.StringArray
}

// GetPosts get posts with pagination
func GetPosts(page int) (result domain.PostListResult) {
	postAccessor := dataaccessor.PostAccessor

	result.TotalCount = postAccessor.GetPostsCount(dataaccessor.PublishedOnly)
	perPageCount := 10
	result.PerPageCount = perPageCount
	offset := perPageCount * (page - 1)
	result.Posts = postAccessor.GetPosts(offset, perPageCount, dataaccessor.PublishedOnly)

	return
}

// GetPost get post with specified id
func GetPost(id uuid.UUID) (post domain.SinglePostItem, err error) {
	db := infrastructure.Db
	sql := `
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
	on t.id = tpa.tag_id
where
	p.id = ?
group by
	p.id
`

	row := db.Raw(sql, id.String()).Row()
	err = row.Scan(&post.CreatedAt, &post.UpdatedAt, &post.ID, &post.UserID, &post.Title, &post.Body, &post.PlainBody, &post.Published, &post.Thumbnail, &post.Tags)

	return
}

// CreatePost create post
func CreatePost(create PostCreate) (postID uuid.UUID, err error) {
	db := infrastructure.Db

	postID = uuid.New()
	post := dataaccessor.Post{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ID:        postID,
		UserID:    create.UserID,
		Title:     create.Title,
		Body:      create.Body,
		PlainBody: create.PlainBody,
		Published: create.Published,
		Thumbnail: create.Thumbnail,
	}

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

	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&post).Error; err != nil {
			return err
		}
		if err = tx.Exec(tagsAttachmentSQL, postID.String(), create.Tags).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		postID = uuid.Nil
	}
	return
}

// UpdatePost updates a post
func (app *PostApplication) UpdatePost(update PostUpdate) (err error) {
	db := infrastructure.Db

	deleteTagsAttachmentSQL := `
delete from
	tags_posts_attachments
where
	post_id = ?
`

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
	tag_name = any(?)`

	err = db.Transaction(func(tx *gorm.DB) error {
		err = db.Exec(deleteTagsAttachmentSQL, update.PostID).Error
		if err != nil {
			return err
		}

		err = db.Exec(updatePostSQL, update.Title, update.Body, update.PlainBody, update.Published, update.Thumbnail, update.PostID, update.UserID).Error
		if err != nil {
			return err
		}

		if err = tx.Exec(tagsAttachmentSQL, update.PostID.String(), update.Tags).Error; err != nil {
			return err
		}

		return nil
	})
	return
}

// DeletePost deletes a post
func (app *PostApplication) DeletePost(userID string, postID string) (err error) {
	db := infrastructure.Db

	postUUID, err := uuid.Parse(postID)
	if err != nil {
		return
	}

	post := Post{
		ID: postUUID,
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		err = tx.Where("post_id = ?", postUUID).Delete(&dataaccessor.TagsPostsAttachments{}).Error
		if err != nil {
			return err
		}

		affected := tx.Delete(&post).RowsAffected
		if affected != 1 {
			return errors.New("No post deleted")
		}

		return nil
	})
	return
}
