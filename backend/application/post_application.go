package application

import (
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/ahmetb/go-linq"
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

// PostGetResult is result of get posts
type PostGetResult struct {
	TotalCount   int
	PerPageCount int
	Posts        []domain.PostListItem
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

// GetAdminPosts get posts with pagination
func (app *PostApplication) GetAdminPosts(page int, userID string, getType domain.GetPostType) (result domain.PostListResult) {
	postAccessor := dataaccessor.PostAccessor

	result.TotalCount = postAccessor.GetAdminPostsCount(getType, userID)
	perPageCount := 10
	result.PerPageCount = perPageCount
	offset := perPageCount * (page - 1)
	result.Posts = postAccessor.GetAdminPosts(offset, perPageCount, getType, userID)

	return
}

// GetPosts get posts with pagination
func GetPosts(page int) (result domain.PostListResult) {
	postAccessor := dataaccessor.PostAccessor

	result.TotalCount = postAccessor.GetPostsCount(domain.Published)
	perPageCount := 10
	result.PerPageCount = perPageCount
	offset := perPageCount * (page - 1)
	result.Posts = postAccessor.GetPosts(offset, perPageCount, domain.Published)

	return
}

// GetPostList get posts
func (app *PostApplication) GetPostList(tags []string, page int, searchWord string, getType domain.GetPostType) (result PostGetResult) {
	perPageCount := 10
	result.PerPageCount = perPageCount
	offset := perPageCount * (page - 1)
	result.Posts = dataaccessor.PostAccessor.GetPostList(tags, searchWord, domain.Published)
	result.TotalCount = len(result.Posts)
	linq.From(result.Posts).Skip(offset).Take(perPageCount).ToSlice(&result.Posts)
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

// CreatePost create a post
func (app *PostApplication) CreatePost(create PostCreate) (postID uuid.UUID, err error) {
	db := infrastructure.Db
	postAccessor := dataaccessor.PostAccessor

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

	err = db.Transaction(func(tx *gorm.DB) error {
		if err := postAccessor.CreatePost(post, tx); err != nil {
			return err
		}
		if err = postAccessor.CreateTagAttachments(postID.String(), create.Tags, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		postID = uuid.Nil
	}
	return

}

// CreatePost create post
func CreatePost(create PostCreate) (postID uuid.UUID, err error) {
	db := infrastructure.Db
	postAccessor := dataaccessor.PostAccessor

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

	err = db.Transaction(func(tx *gorm.DB) error {
		if err := postAccessor.CreatePost(post, tx); err != nil {
			return err
		}
		if err = postAccessor.CreateTagAttachments(postID.String(), create.Tags, tx); err != nil {
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
	err = db.Transaction(func(tx *gorm.DB) error {
		postAccessor := dataaccessor.PostAccessor

		err = postAccessor.DeleteAttachedTags(update.PostID, tx)
		if err != nil {
			return err
		}

		err = postAccessor.UpdatePost(update.PostID, update.Title, update.Body, update.PlainBody, update.Published, update.Thumbnail, update.UserID, tx)
		if err != nil {
			return err
		}

		err = postAccessor.CreateTagAttachments(update.PostID.String(), update.Tags, tx)
		if err != nil {
			return err
		}

		return nil
	})
	return
}

// DeletePost deletes a post
func (app *PostApplication) DeletePost(userID string, postID string) (err error) {
	db := infrastructure.Db
	postAccessor := dataaccessor.PostAccessor

	postUUID, err := uuid.Parse(postID)
	if err != nil {
		return
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		err = postAccessor.DeleteAttachedTags(postUUID, tx)
		if err != nil {
			return err
		}

		err = postAccessor.DeletePost(postUUID, tx)
		if err != nil {
			return err
		}

		return nil
	})
	return
}
