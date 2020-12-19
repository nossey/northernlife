package dataaccessor

import (
	"github.com/google/uuid"
)

// TagsPostsAttachments is attachment relation of posts and tags
type TagsPostsAttachments struct {
	PostID uuid.UUID `gorm:"post_id"`
}
